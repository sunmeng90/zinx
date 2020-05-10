package main

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"github.com/sunmeng90/zinx/znet"
	"strconv"
)

func main() {
	server := znet.NewServer("zinx")
	server.SetOnConnStart(func(conn ziface.IConn) {
		fmt.Println("connection ", conn.GetConnID(), " established")
		//TODO: the send message on start is displayed and dismissed immediately
		if err := conn.SendMsg(1, []byte("on connection established from server")); err != nil {
			fmt.Println("can't send message to client for establish signal", err)
		}
		fmt.Println("set property on connection")
		conn.SetProp("name", "connection"+strconv.Itoa(int(conn.GetConnID())))
	})
	server.SetOnConnStop(func(conn ziface.IConn) {
		// can't send message here, connection maybe closed, before the message is processed by the worker
		fmt.Println("connection ", conn.GetConnID(), " stopped")

	})
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &EchoRouter{})
	server.Serve()
}

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("handle ping request msg id: ", req.MsgId(), " data: ", string(req.Data()))
	err := req.Conn().SendMsg(100+req.MsgId(), []byte(".......pong......."))
	if err != nil {
		fmt.Println("failed to write data to client", err)
	}
}

type EchoRouter struct {
	znet.BaseRouter
}

func (p *EchoRouter) Handle(req ziface.IRequest) {
	fmt.Println("handle echo request msg id: ", req.MsgId(), " data: ", string(req.Data()))
	err := req.Conn().SendMsg(200+req.MsgId(), []byte("......."+string(req.Data())+"......."))
	if err != nil {
		fmt.Println("failed to write data to client", err)
	}
}
