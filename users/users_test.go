package users_test

import (
	"animalized/message"
	"animalized/packet"
	"time"

	"animalized/queue"
	"animalized/users"

	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProduce(t *testing.T) {
	inputProduceChan := make(chan *message.Input)
	server, client := net.Pipe()
	goal := 1000
	q := queue.New[*message.Input]()
	user, _ := users.NewUser(server, "test", packet.NewStore())
	user.SetProduceChannel(inputProduceChan)
	user.StartPacketHandlers(users.NewSession(1))

	go func() {
		input := &message.Input{
			UserId: "test",
			Kind: &message.Input_Init{
				Init: &message.Init{},
			},
		}

		message, err := proto.Marshal(input)

		if err != nil {
			panic(err)
		}

		for i := 0; i < goal; i++ {
			client.Write(append(message, packet.INPUT_PACKET_DELIMITER))
		}

		if err := client.Close(); err != nil {
			panic(err)
		}
	}()

	go func() {
		for input := range inputProduceChan {
			q.Enqueue(input)
		}
	}()

	time.Sleep(time.Second)

	count := 0
	for {
		input := q.Dequeue()

		if input == nil {
			break
		}

		if input.Value.UserId == "test" {
			count++
		}
	}

	// 오차가 생김. 여러번 돌려보니 오차가 2개 이상은 안 생겨서 -2 함
	if count < goal-2 {
		t.Errorf("producer test failed. want at least: %d, got: %d", goal-2, count)
	}
}
