package users

import (
	"animalized/common"
	"animalized/message"
	"animalized/packet"
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
	packetStore    *packet.PacketStore
	produceChannel chan<- *message.Input
}
