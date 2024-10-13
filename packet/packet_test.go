package packet_test

import (
	"animalized/message"
	"animalized/packet"
	"net"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func TestInputParsing(t *testing.T) {
	server, client := net.Pipe()
	direction := int32(1)

	tcs := []struct {
		desc  string
		input *message.Input
	}{
		{
			desc: "parse single packet",
			input: &message.Input{
				Type:   1,
				UserId: "test",
			},
		},
		{
			desc: "parse optional field packet",
			input: &message.Input{
				Type:      1,
				UserId:    "test",
				Direction: &direction,
			},
		},
		{
			desc: "parse type 2 packet",
			input: &message.Input{
				Type:   2,
				UserId: "test",
			},
		},
	}

	for _, tc := range tcs {
		go func() {
			message, _ := proto.Marshal(tc.input)
			client.SetWriteDeadline(time.Now().Add(2 * time.Second))
			client.Write(append(message, '$'))
		}()

		input, err := packet.ParseInput(&server)

		if err != nil {
			t.Fatal(err)
		}

		if input.UserId != "test" {
			t.Fatal("user id not matched")
		}
	}
}
