package main

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"github.com/sunmeng90/zinx/znet"
)

func main() {
	server := znet.NewServer("zinx")
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
