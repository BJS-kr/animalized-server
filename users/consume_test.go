package users_test

import (
	"animalized/common"

	"animalized/message"
	"animalized/packet"
	"animalized/queue"
	"animalized/users"
	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestConsume(t *testing.T) {
	server, client := net.Pipe()
	q := queue.New[*message.Input]()
	q.Enqueue(&message.Input{
		Type:   1,
		UserId: "test",
	})

	u := &users.User{
		Id:   "test",
		Conn: server,
		Base: common.Base{
			Inputs: q,
		},
	}

	input := new(message.Input)

	go func() {
		users.Consume(u)
	}()

	buf := make([]byte, packet.BUFFER_SIZE)
	client.Read(buf)
	proto.Unmarshal(buf[:len(buf)-1], input)

	if input.UserId != "test" {
		t.Errorf("expected value not matched. want: %s, actual: %s", "test", input.UserId)
	}
}