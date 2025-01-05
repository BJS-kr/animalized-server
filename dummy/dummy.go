package dummy

import (
	"animalized/message"
	"encoding/binary"
	"net"

	"google.golang.org/protobuf/proto"
)

func Dummy(userId string) {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9988")
	conn, _ := net.DialTCP("tcp", nil, addr)

	sendInit(conn, userId)
}

func send(conn net.Conn, message proto.Message) {
	data, _ := proto.Marshal(message)
	lenBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(lenBuf, uint16(len(data)))
	conn.Write(lenBuf)
	conn.Write(data)
}

func sendInit(conn net.Conn, userId string) {
	initMessage := &message.Input{
		UserId: userId,
		Kind:   &message.Input_Init{},
	}

	send(conn, initMessage)
}

func sendCreateRoom(conn net.Conn, roomName string) {
	createRoomMessage := &message.Input{
		Kind: &message.Input_Lobby{
			Lobby: &message.Lobby{
				Type:     message.Lobby_CREATE_ROOM,
				RoomName: roomName,
				MaxUsers: 6,
			},
		},
	}

	send(conn, createRoomMessage)
}

func sendJoinRoom(conn net.Conn, roomName string) {
	joinRoomMessage := &message.Input{
		Kind: &message.Input_Lobby{
			Lobby: &message.Lobby{
				Type:     message.Lobby_JOIN_ROOM,
				RoomName: roomName,
			},
		},
	}

	send(conn, joinRoomMessage)
}

func sendLeaveRoom(conn net.Conn, roomName string) {
	leaveRoomMessage := &message.Input{
		Kind: &message.Input_Room{
			Room: &message.Room{
				Type:     message.Room_QUIT,
				RoomName: roomName,
			},
		},
	}

	send(conn, leaveRoomMessage)
}

func sendStartGame(conn net.Conn, roomName string) {
	startGameMessage := &message.Input{
		Kind: &message.Input_Room{
			Room: &message.Room{
				Type:     message.Room_START,
				RoomName: roomName,
			},
		},
	}

	send(conn, startGameMessage)
}
