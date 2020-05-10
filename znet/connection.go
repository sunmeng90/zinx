package znet

import (
	"errors"
	"fmt"
	"github.com/sunmeng90/zinx/utils"
	"github.com/sunmeng90/zinx/ziface"
	"io"
	"net"
	"sync"
)

type Conn struct {
	Server ziface.IServer

	// connection socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	MsgChan chan []byte

	ExitChan chan bool

	MessageHandle ziface.IMessageHandle

	props map[string]interface{}

	propsLock sync.RWMutex
}

// There is no such thing as a "pointer to an interface" (technically, you can use one, but generally you don't need it).
// https://www.howtobuildsoftware.com/index.php/how-do/bKOd/pointers-struct-interface-casting-go-cast-a-struct-pointer-to-interface-pointer-in-golang
// TODO: can't pass server struct pointer to server *ziface.IServer
func NewConn(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandle ziface.IMessageHandle) *Conn {
	c := &Conn{
		Server:        server, // TODO: it's wired, refactor this
		Conn:          conn,
		ConnID:        connID,
		isClosed:      false,
		MsgChan:       make(chan []byte),
		ExitChan:      make(chan bool, 1),
		MessageHandle: msgHandle,
		props:         make(map[string]interface{}),
		propsLock:     sync.RWMutex{},
	}
	server.ConnManager().Add(c)
	return c
}

func (c *Conn) Start() {
	fmt.Println("Start conn ", c.ConnID)
	go c.StartReader()
	go c.StartWriter()
	c.Server.CallOnConnStart(c)
}

func (c *Conn) StartReader() {
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
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MessageHandle.SendReqToQueue(&req)
		} else {
			go c.MessageHandle.DoMsgHandler(&req)
		}
	}
}

func (c *Conn) StartWriter() {
	fmt.Println("writer goroutine")
	defer fmt.Println("writer exit for remote", c.RemoteAddr())
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("writer send data error", err)
				return
			}
		case <-c.ExitChan:
			fmt.Println("writer got signal to exit")
			return
		}
	}
}

func (c *Conn) Stop() {
	fmt.Println("Stop conn ", c.ConnID)

	if c.isClosed {
		return
	}
	c.Server.CallOnConnStop(c)
	c.Conn.Close()
	c.ExitChan <- true
	c.isClosed = true
	c.Server.ConnManager().Remove(c)
	close(c.ExitChan)
	close(c.MsgChan)
}

func (c *Conn) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Conn) GetConnID() uint32 {
	return c.ConnID
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
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
	c.MsgChan <- msgBytes
	return nil
}

func (c *Conn) SetProp(key string, val interface{}) {
	c.propsLock.Lock()
	defer c.propsLock.Unlock()
	c.props[key] = val
}

func (c *Conn) Prop(key string) (val interface{}, err error) {
	c.propsLock.RLock()
	defer c.propsLock.RUnlock()
	val, ok := c.props[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return val, nil
}

func (c *Conn) RemoveProp(key string) {
	c.propsLock.Lock()
	defer c.propsLock.Unlock()
	delete(c.props, key)
}
