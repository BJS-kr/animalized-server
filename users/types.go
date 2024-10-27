package users

import (
	"animalized/common"
	"animalized/message"
	"net"
	"sync"
)

type DistributableUsers struct {
	common.Distributable
	*Users
}

type Users struct {
	mtx  sync.RWMutex
	Max  int
	list []*User
}

type User struct {
	common.Distributable
	Conn           net.Conn
	Id             string
	produceChannel chan<- *message.Input
}
