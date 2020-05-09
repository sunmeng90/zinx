package main

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"github.com/sunmeng90/zinx/znet"
)

func main() {
	server := znet.NewServer("zinx")
	server.AddRouter(&PingRouter{})
	server.Serve()
}

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("handle ping request msg id: ", req.MsgId(), " data: ", string(req.Data()))
	err := req.Conn().SendMsg(req.MsgId(), []byte(".......pong......."))
	if err != nil {
		fmt.Println("failed to write data to client", err)
	}
}
