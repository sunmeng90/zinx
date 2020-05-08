package znet

import (
	"github.com/sunmeng90/zinx/ziface"
	"net"
)

type Conn struct {
	// connection socket
	Conn *net.TCPConn

	//
	ConnID uint32

	isClosed bool

	handleAPI ziface.HandleFun

	ExitChan chan bool
}

func NewConn(conn *net.TCPConn, connID uint32, callback ziface.HandleFun) *Conn {
	return &Conn{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback,
		ExitChan:  make(chan bool, 1),
	}
}
