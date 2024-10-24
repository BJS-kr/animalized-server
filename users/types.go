package users

import (
	"animalized/common"
	"net"
	"sync"
)

type DistributableUsers struct {
	common.Base
	Users *Users
}

type Users struct {
	mtx  sync.RWMutex
	Max  int
	list []*User
}

type User struct {
	common.Base
	Conn net.Conn
	Id   string
}

type OutgoingHandler func(*User)
