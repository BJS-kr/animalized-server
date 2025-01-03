package controller

import (
	"animalized/lobby"
	"animalized/message"
	"animalized/rooms"
)

type Controller struct {
	Lobby *lobby.Lobby
	Rooms *rooms.Rooms
}

func New(maxUsers int) *Controller {
	c := new(Controller)

	c.Lobby = lobby.New(maxUsers)
	c.Rooms = rooms.New()

	c.Lobby.StartStreaming(c.lobbyHandler, LOBBY_TICK_RATE)

	return c
}

/**
 * CAUTION: Use as System Direct Input
 */
func (c *Controller) MakeLobbyState(userId string) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_Lobby{
		Lobby: &message.Lobby{
			Type: message.Lobby_STATE,
		},
	}

	return input
}

/**
 * CAUTION: Use as System Direct Input
 */
func (c *Controller) MakeJoinInput(userId string, roomName string) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_Lobby{
		Lobby: &message.Lobby{
			Type:     message.Lobby_JOIN_ROOM,
			RoomName: roomName,
		},
	}

	return input
}

/**
 * CAUTION: Use as System Input
 */
func (c *Controller) MakeRoomStateInput(userId string, roomName string) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_Room{
		Room: &message.Room{
			Type:     message.Room_STATE,
			RoomName: roomName,
		},
	}

	return input
}

/**
 * CAUTION: Use as System Direct Input
 */
func (c *Controller) MakeRoomStateDirectInput(userId string, roomName string, room *rooms.Room) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_Room{
		Room: &message.Room{
			Type:      message.Room_STATE,
			RoomName:  roomName,
			RoomState: room.MakeRoomState(roomName),
		},
	}

	return input
}

/**
 * CAUTION: Use as System Direct Input
 */
func (c *Controller) MakeQuitRoomInput(userId string, roomName string) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_Lobby{
		Lobby: &message.Lobby{
			Type:     message.Lobby_QUIT_ROOM,
			RoomName: roomName,
		},
	}

	return input
}

/**
 * CAUTION: Use as System Direct Input
 */
func (c *Controller) MakeGameStartInput(userId string, roomName string, userCharacterTypes rooms.UserCharacterTypes, terrains []*message.Terrain) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_Room{
		Room: &message.Room{
			Type:               message.Room_START,
			RoomName:           roomName,
			UserCharacterTypes: userCharacterTypes,
			Terrains:           terrains,
		},
	}

	return input
}

func (c *Controller) MakeGameResultInput(userId string, roomName string) *message.Input {
	input := new(message.Input)

	input.UserId = userId
	input.Kind = &message.Input_GameResult{
		GameResult: &message.GameResult{
			WinnerId: userId,
			RoomName: roomName,
		},
	}

	return input
}
