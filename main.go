package main

import (
	"animalized/message"
	"animalized/packet"
	"errors"
	"log/slog"
	"net"

	"time"
)

const (
	READ_DEADLINE = time.Duration(time.Minute)
)

const (
	INIT = iota + 1
	MOVE
	ATTACK
)

var userConns []*net.TCPConn = make([]*net.TCPConn, 0)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9988")

	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		panic(err)
	}

	inputProduceChannel := make(chan *message.Input, 100)

	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		launchInputProducer(conn, inputProduceChannel)
	}
}

func SetTimeLimit(conn *net.TCPConn) error {
	if err := conn.SetReadDeadline(time.Now().Add(READ_DEADLINE)); err != nil {
		return err
	}

	// total 3min
	// idle time default 15sec + interval default 15sec * 9
	// https://pkg.go.dev/net#KeepAliveConfig
	if err := conn.SetKeepAlive(true); err != nil {
		return err
	}

	return nil
}

func launchInputProducer(conn *net.TCPConn, inputProduceChannel chan<- *message.Input) {
	if err := Initialize(conn); err != nil {
		slog.Error(err.Error())
		return
	}

	if err := SetTimeLimit(conn); err != nil {
		slog.Error(err.Error())
		return
	}

	userConns = append(userConns, conn)

	go ProduceInput(conn, inputProduceChannel)
}

func Initialize(conn *net.TCPConn) error {
	initInput, err := packet.ParseInput(conn)

	if err != nil {
		return err
	}

	if !IsInitPacket(initInput) {
		return errors.New("init packet type invalid")
	}

	return nil
}

func ProduceInput(conn *net.TCPConn, inputProduceChannel chan<- *message.Input) {
	for {
		input, err := packet.ParseInput(conn)

		if err != nil {
			return
		}

		inputProduceChannel <- input

		if err := conn.SetReadDeadline(time.Now().Add(READ_DEADLINE)); err != nil {
			return
		}
	}
}

// 연결된 모든 TCP 커넥션을 순회하며 input을 뿌리는 역할
// 뿌린다: 그 커넥션이 보내야 할 대기열에 넣는다.
// 여기서 가장 고민인 점은 rw의 모순적인 상황이다.
// 고루틴 여러개가 데이터를 삽입하는 것은 순차적으로 락 없이 가능하지만
// 그걸 읽어서 브로드캐스트할 때는 어떻게 해야할까?
// 브로드캐스트 된 지점까지는 보내야하는 데이터에서 빠져야한다.
// 그러면 insert-delete가 동시에 일어나야한다는 것인데, 방안은 크게 세가지다.
// 1. 그냥 틱당 락을 걸어서 브로드캐스팅 한 다음 보낸 부분까지 제외한 다음 insert를 받는다. 가장 naive함
// 2. 보낸 부분의 위치를 기억한다. 예를 들어 보낸 부분의 index가 1000이 넘어가면 그때 락을 걸고 인풋 큐를 초기화한다. 그러나 이 방법은 정확히 틱에 맞춰 보낸다는 보장이 없다.
// lock을 얻기까지 얼마나 많은 선행 lock이 존재할지 정확히 예측할 수 없을 뿐더러, 즉시 락을 얻는다고 하더라도 전파 작업시간이 틱을 초과할 수도 있다.
// 1,2의 가장 큰 문제는 다름아닌 쓰기에서 락이 걸린다는 것이다. producer들도 언제 consumer가 큐를 초기화할지 알 수 없으니 매번 락을 걸어야 정합성이 보장된다.
// 3번째 방법은 lock-free를 쓰는 것이다. 사실 LL만 써도 lock을 안 걸어도 된다.
// 3-1. LL
// 3-2. sync.Map
// 아무래도 sync.Map이 좋아보인다. range로 간편하게 돌 수 있을 뿐더러, key를 input객체 메모리 주소로하고 val은 그냥 nil로 넣은 후(그래야 key가 안겹치니까. uuid같은거 새로 만들기 싫음)
// Range를 돌면서 cb안에서 LoadAndStore로 전파하면 될 듯
// 단, 약점은 sync.Map자체가 이미 mutex를 내장하고 있다는 점
// sync.Map이 퍼포먼스가 안좋다는 글도 있다.
// 핵심 요인. sync.Map.Range는 삽입한 순서대로 돌지 않는다고 한다.
// 3.1로 가자
// 사고의 흐름: mutex & slice -> sync.Map -> LL -> lock free LL
// mutex & slice를 사용하지 않는 이유: r & w가 intensive한데 거기에 mutex를 매번 거는 것은 낭비같다.
// sync.Map을 선택하지 않는 이유: 퍼포먼스, 내부적 mutex문제도 있지만 핵심적으로 순서대로 range하는 보장이 없다.
// LL을 선택하지 않는 이유: actor와 dispatcher가 LL을 사용하게 되는데, actor는 tail, dispatcher는 head만 사용하니까 일견 문제가 없어보이지만
// 문제는 큐가 끝에 도달할때이다. head와 tail에 actor와 dispatcher가 동시에 접근하게 될 가능성이 있고, 현재 위치가 어디인지 모르니 안정성을 위해서 매번 lock을 걸어야한다.
// lock free LL을 고민하고 선택한 이유: 일단, 구현자체가 눈에 잘 들어오지 않는다. lock free에 이해도 깊지 않다.
// 그러나 결국 접근하려는 스레드가 두 개(actor, dispatcher)이므로 lock-contention이 높지 않고, rw rate가 높기 때문에 lock을 매번 거느니 lock free가 낫다고 판단했다.
// dispatcher와 워커 간에 buffered 채널을 사용할 수 없는 이유: dispatcher가 block될 위험 존재. dispatcher는 워커의 상태에 영향을 받지 않아야 한다.
// 그렇다면 worker는 dispatcher로 부터 받은 인풋을 바로 수령하고 자신이 컨트롤하는 하는 큐에 넣으면 되는 것 아닌가 할 수 있지만,
// 그렇게 되면 어차피 1스레드가 전송과 채널 수령을 모두 해야하므로 마찬가지로 전송에 지연이 생겨 채널이 막힐 위험이 있고,
// 워커 자체가 복수의 고루틴으로 구성된다고 해도 그렇다면 구현의 복잡성이 올라갈 뿐 아니라 대상큐에 락이 걸리거나 워커간 채널이 병목될 위험은 여전히 존재하게 된다.
// 결론적으로 dispatcher와 워커 간에도 lock free queue를 사용해야하고, 대기열이 비정상적인 수치까지 도달하면 워커와 커넥션 자체를 종료시켜야한다.
// 추가: dispatcher가 직접 connection을 돌며 전파하지 않는이유: 워커에 뿌려주는 것과 직접 conn을 돌면서 보내는 것은 지연 관리에 큰 영향이 있다.
// 10개의 클라이언트가 접속했다고 가정해보면 10개를 모두 기다리거나,
// 기다리지 않기 위해선 10개의 고루틴을 띄우고 다음 루프로 진입해야 하는데 아무리 goroutine이 경량이라지만
// 빠른 속도로 spin을 도는 상황에서 매번 복수의 goroutine을 띄울수도 없거니와,
// 유저가 10이 아니라 100명, 1000명이라고 생각해보면 문제는 더욱더 심해진다.
func Broker(inputProduceChannel <-chan *message.Input) {
	for input := range inputProduceChannel {

	}
}

// TCP 연결과 동시에 첫 패킷은 무조건 유저의 아이디를 담은 패킷이어야 한다.
// 일정시간 동안 패킷이 도착하지 않거나 첫 패킷이 아이디가 아니라면 커넥션은 그대로 버린다.
// 아이디 패킷이 정상적으로 도착할 경우 커넥션을 저장한다.
func IsInitPacket(input *message.Input) bool {
	return input.GetType() == INIT
}
func StoreConn() {

}

func Propagate() {

}

// 메인 루틴으로부터 받은 인풋들을 보낸다. 고루틴으로서 큐에 인풋이 있으면 무조건 보낸다.
// 이것만은 간단 구현이라도 순차 작업 할 수 없다. 인덱스의 뒤에 있는 유저가 큰 손해를 보게 되게 때문.
// io작업이기 떄문에 단순히 append와는 다르다.
// 받은 인풋을 또 다시 수동으로 따로 큐를 만드는 대신 그냥 buffered channel로 한다.
func WriteInputInOrder() {

}
