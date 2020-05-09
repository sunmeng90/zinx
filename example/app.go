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

func (p *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("pre handle ping ")
	conn := req.Conn()
	if _, err := conn.GetTCPConn().Write([]byte("pre handle")); err != nil {
		fmt.Println("pre handle ping error", err)
	}
}

func (p *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("handle ping ")
	conn := req.Conn()
	if _, err := conn.GetTCPConn().Write([]byte("......ping......")); err != nil {
		fmt.Println("handle ping error", err)
	}
}

func (p *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("post handle ping ")
	conn := req.Conn()
	if _, err := conn.GetTCPConn().Write([]byte("post handle")); err != nil {
		fmt.Println("post handle ping error", err)
	}
}
