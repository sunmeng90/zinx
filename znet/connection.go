package znet

import (
	"errors"
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"io"
	"net"
)

type Conn struct {
	// connection socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChan chan bool

	MessageHandle ziface.IMessageHandle
}

func NewConn(conn *net.TCPConn, connID uint32, msgHandle ziface.IMessageHandle) *Conn {
	return &Conn{
		Conn:          conn,
		ConnID:        connID,
		isClosed:      false,
		ExitChan:      make(chan bool, 1),
		MessageHandle: msgHandle,
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
		dp := NewDataPack()
		headData := make([]byte, dp.HeadLen())
		// don't read all bytes in connection, we need read head and get data len, and only consume
		// what's meaningful for current packet
		_, err := io.ReadFull(c.Conn, headData)
		if err != nil {
			fmt.Println("failed to read head data", err)
			break
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("failed to unpack head data")
			break
		}
		data := make([]byte, msg.Len())
		_, err = io.ReadFull(c.Conn, data)
		if err != nil {
			fmt.Println("failed to read data", err)
			break
		}
		msg.SetData(data)
		req := Request{
			conn: c,
			msg:  msg,
		}
		go c.MessageHandle.DoMsgHandler(&req)
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

func (c *Conn) SendMsg(id uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection is closed")
	}
	// pack message
	msgBytes, err := NewDataPack().Pack(NewMsg(id, data))
	if err != nil {
		return err
	}
	c.Conn.Write(msgBytes)
	return nil
}
