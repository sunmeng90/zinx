package main

import "github.com/sunmeng90/zinx/znet"

func main() {
	server := znet.NewServer("zinx")
	server.Serve()
}
