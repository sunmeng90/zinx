package znet

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"net"
)

type Conn struct {
	// connection socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChan chan bool

	Router ziface.IRouter
}

func NewConn(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Conn {
	return &Conn{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
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
		if _, err := c.Conn.Read(buf); err != nil {
			fmt.Println("Failed to read from conn")
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}
		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)

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
	return c.Conn
}

func (c *Conn) GetConnID() uint32 {
	return c.ConnID
}

func (c *Conn) RemoteAddr() net.Addr {
	panic("implement me")
}

func (c *Conn) Send(data []byte) error {
	panic("implement me")
}
