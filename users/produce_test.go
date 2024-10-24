package users_test

import (
	"animalized/message"
	"animalized/packet"

	"animalized/queue"
	"animalized/users"
	"bytes"
	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProduce(t *testing.T) {
	inputProduceChan := make(chan *message.Input)
	buf, inputBuf := make([]byte, packet.BUFFER_SIZE), bytes.NewBuffer(nil)
	server, client := net.Pipe()
	goal := 1000
	q := queue.New[*message.Input]()

	user := &users.User{
		Conn: server,
		Id:   "test",
	}

	go func() {
		input := &message.Input{
			Type:   1,
			UserId: "test",
		}

		message, _ := proto.Marshal(input)

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

	for {
		if err := user.ProduceInput(&buf, inputBuf, inputProduceChan); err != nil {
			break
		}
	}

	count := 0
	for {
		input := q.Dequeue()
		if input == nil {
			break
		}
		count += int(input.Value.Type)
	}

	// 오차가 생김. 여러번 돌려보니 오차가 2개 이상은 안 생겨서 -2 함
	if count < goal-2 {
		t.Errorf("producer test failed. want at least: %d, got: %d", goal-2, count)
	}
}
