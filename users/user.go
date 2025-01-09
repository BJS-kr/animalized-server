package users

import (
	"animalized/common"
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"bytes"
	"encoding/binary"
	"errors"
	"log/slog"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

type User struct {
	Conn           net.Conn
	Id             string
	packetStore    *packet.PacketStore
	produceChannel chan<- *message.Input
	outgoingQueue  *queue.Queue[*message.Input]
	sizeBuf        []byte
	Tick           chan common.Signal
}

var ErrsUserIdNotMatched = errors.New("user id not matched")

func NewUser(conn net.Conn, id string, packetStore *packet.PacketStore) (*User, error) {
	userIdLen := len(id)
	if userIdLen == 0 || userIdLen > 10 {
		return nil, errors.New("empty or longer than 10 length id not allowed")
	}
	u := new(User)
	u.outgoingQueue = queue.New[*message.Input]()
	u.Id = id
	u.Conn = conn
	u.packetStore = packetStore
	u.Tick = make(chan common.Signal)
	u.sizeBuf = make([]byte, 2)

	return u, nil
}

func (u *User) SetProduceChannel(ch chan<- *message.Input) {
	u.produceChannel = ch
}

// 유저로부터 수집된 인풋들을 중계 스택으로 쌓는다.
// 복수형인 이유는 패킷파싱할 때 일단 커넥션 타고 있는 것들은 싹 순회하고 넣을 예정이라서
// 일단 지금 생각으로는 buffered channel로 넣으면 될 것 같다.
// 패킷 타입은 하나로 통일한다. 로직을 간단화하고 시간순서 맞추기도 편하다.
func (u *User) ProduceInput() (*message.Input, error) {
	input, err := u.packetStore.ParseInput(u.Conn)

	if err != nil {
		return nil, err
	}

	err = u.validateInput(input)

	if err != nil {
		return nil, err
	}

	if err := u.Conn.SetReadDeadline(time.Now().Add(READ_DEADLINE)); err != nil {
		return nil, err
	}

	return input, nil
}

func (u *User) validateInput(input *message.Input) error {
	if input.UserId != u.Id {
		return ErrsUserIdNotMatched
	}

	return nil
}

func (u *User) StartPacketHandlers(session *Session) {
	go u.handleIncoming(session)
	go u.handleOutgoing()
}

func (u *User) handleIncoming(session *Session) {
	for {
		input, err := u.ProduceInput()

		if err != nil {
			slog.Error(err.Error())
			if errors.Is(err, ErrsUserIdNotMatched) {
				continue
			}

			session.Quit(u)
			u.Conn.Close()
			return
		}

		u.produceChannel <- input
	}
}

func (u *User) handleOutgoing() {
	for range u.Tick {
		for {
			n := u.outgoingQueue.Dequeue()

			if n == nil || n.Value == nil {
				break
			}

			msg, err := proto.Marshal(n.Value)

			if err != nil {
				slog.Error(err.Error())
				continue
			}
			if len(msg) > packet.BUFFER_SIZE {
				slog.Error("length of packet size must not exceed BUFFER_SIZE")
			}

			binary.BigEndian.PutUint16(u.sizeBuf, uint16(len(msg)))

			_, err = u.Conn.Write(bytes.Join([][]byte{u.sizeBuf, msg}, nil))

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			message.Pool.Put(n.Value)
		}
	}
}
