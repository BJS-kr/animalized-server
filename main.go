package main

import (
	"log/slog"
	"net"
	"time"
)

const READ_DEADLINE = time.Duration(time.Minute)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9988")

	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			slog.Error(err.Error())
			continue
		}

		handler(conn)
	}
}

func handler(conn *net.TCPConn) {
	if err := SetTimeLimit(conn); err != nil {
		slog.Error(err.Error())
		return
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

// TCP 연결과 동시에 첫 패킷은 무조건 유저의 아이디를 담은 패킷이어야 한다.
// 일정시간 동안 패킷이 도착하지 않거나 첫 패킷이 아이디가 아니라면 커넥션은 그대로 버린다.
// 아이디 패킷이 정상적으로 도착할 경우 커넥션을 저장한다.
func IsInitPacket() {

}
func StoreConn() {

}

// 유저로부터 수집된 인풋들을 중계 스택으로 쌓는다.
// 복수형인 이유는 패킷파싱할 때 일단 커넥션 타고 있는 것들은 싹 순회하고 넣을 예정이라서
// 일단 지금 생각으로는 buffered channel로 넣으면 될 것 같다.
// 패킷 타입은 하나로 통일한다. 로직을 간단화하고 시간순서 맞추기도 편하다.
func ParsePacket() {

}
func CollectInputs() {

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
func Propagate() {

}

// 메인 루틴으로부터 받은 인풋들을 보낸다. 고루틴으로서 큐에 인풋이 있으면 무조건 보낸다.
// 이것만은 간단 구현이라도 순차 작업 할 수 없다. 인덱스의 뒤에 있는 유저가 큰 손해를 보게 되게 때문.
// io작업이기 떄문에 단순히 append와는 다르다.
// 받은 인풋을 또 다시 수동으로 따로 큐를 만드는 대신 그냥 buffered channel로 한다.
func WriteInputInOrder() {

}
