package ziface

import "net"

type IConn interface {
	Start()
	Stop()
	// socket conn
	GetTCPConn() *net.TCPConn
	GetConnID() uint32
	// client address
	RemoteAddr() net.Addr

	SendMsg(id uint32, data []byte) error

	SetProp(key string, val interface{})

	Prop(key string) (val interface{}, err error)

	RemoveProp(key string)
}

type HandleFun func(*net.TCPConn, []byte, int) error
