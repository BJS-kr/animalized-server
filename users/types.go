package users

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
	Quit       chan struct{}
}

type Users struct {
	mtx   sync.RWMutex
	Max   int
	users []*User
}
