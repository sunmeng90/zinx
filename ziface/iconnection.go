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

	Send(data []byte) error
}

type HandleFun func(*net.TCPConn, []byte, int) error
