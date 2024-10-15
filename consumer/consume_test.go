package consumer_test

import (
	"animalized/consumer"
	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"animalized/user"
	"net"
	"sync"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestConsume(t *testing.T) {
	var wg sync.WaitGroup
	server, client := net.Pipe()
	q := queue.New[*message.Input]()
	q.Enqueue(&message.Input{
		Type:   1,
		UserId: "test",
	})

	u := &user.User{
		Id:         "test",
		Conn:       server,
		InputQueue: q,
	}

	input := new(message.Input)

	wg.Add(1)
	go func() {
		buf := make([]byte, packet.BUFFER_SIZE)
		for {
			size, _ := client.Read(buf)

			if size > 0 {
				proto.Unmarshal(buf[:len(buf)-1], input)
				wg.Done()
				return
			}
		}
	}()

	consumer.Consume(u)
	wg.Wait()

	if input.UserId != "test" {
		t.Errorf("expected value not matched. want: %s, actual: %s", "test", input.UserId)
	}
}
