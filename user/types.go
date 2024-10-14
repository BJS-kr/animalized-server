package user

import (
	"animalized/message"
	"animalized/queue"
	"net"
	"sync"
)

type User struct {
	Conn       net.Conn
	InputQueue *queue.Queue[*message.Input]
	Id         string
}

type Users struct {
	mtx   sync.RWMutex
	users []*User
}
