package packet_test

import (
	"animalized/message"
	"animalized/packet"
	"net"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestInputParsing(t *testing.T) {
	server, client := net.Pipe()

	tcs := []struct {
		desc  string
		input *message.Input
	}{
		{
			desc: "parse init packet",
			input: &message.Input{
				UserId: "test",
				Kind:   &message.Input_Init{},
			},
		},
		{
			desc: "parse move packet",
			input: &message.Input{

				UserId: "test",
				Kind: &message.Input_Operation{
					Operation: &message.Operation{
						Type:      message.Operation_MOVE,
						Direction: message.Operation_UP,
					},
				},
			},
		},
		{
			desc: "parse lobby packet",
			input: &message.Input{
				UserId: "test",
				Kind: &message.Input_Lobby{
					Lobby: &message.Lobby{
						Type:     message.Lobby_CREATE,
						RoomName: "test",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		go func() {
			message, _ := proto.Marshal(tc.input)
			client.Write(append(message, packet.INPUT_PACKET_DELIMITER))
		}()

		packetStore := packet.NewStore()
		input, err := packetStore.ParseInput(server)

		if err != nil {
			t.Fatal(err)
		}

		if input.UserId != "test" {
			t.Fatal("user id not matched")
		}
	}
}
