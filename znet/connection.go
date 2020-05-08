package znet

import (
	"fmt"
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

func (c *Conn) Start() {
	fmt.Println("Start conn ", c.ConnID)
	go c.StartRead()
}

func (c *Conn) StartRead() {
	fmt.Println("Start reading")
	defer fmt.Println("Stop read on conn ", c.ConnID)
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read from conn")
			continue
		}
		// pass bytes to handler
		if err = c.handleAPI(c.Conn, buf, n); err != nil {
			fmt.Println("Handle error", err)
		}
	}
}

func (c *Conn) Stop() {
	fmt.Println("Stop conn ", c.ConnID)
	if c.isClosed {
		return
	}
	c.Conn.Close()
	c.isClosed = true
	close(c.ExitChan)
}

func (c *Conn) GetTCPConn() *net.TCPConn {
	panic("implement me")
}

func (c *Conn) GetConnID() uint32 {
	panic("implement me")
}

func (c *Conn) RemoteAddr() net.Addr {
	panic("implement me")
}

func (c *Conn) Send(data []byte) error {
	panic("implement me")
}
