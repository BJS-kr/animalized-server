package users

import (
	"animalized/message"
	"animalized/queue"
	"net"
	"sync"
)

type User struct {
	Conn   net.Conn
	Inputs *queue.Queue[*message.Input]
	Id     string
	Stop   chan struct{}
}

type Users struct {
	mtx   sync.RWMutex
	Max   int
	users []*User
}
