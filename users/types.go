package users

import (
	"animalized/common"
	"net"
	"sync"
)

type DistributableUsers struct {
	common.Distributable
	Users *Users
}

type Users struct {
	mtx  sync.RWMutex
	Max  int
	list []*User
}

type User struct {
	common.Distributable
	Conn net.Conn
	Id   string
}

type OutgoingHandler func(*User)
