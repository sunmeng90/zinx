package main

import (
	"fmt"
	"github.com/sunmeng90/zinx/znet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Client start")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Print("failed to connect to server ", err)
		return
	}
	for {
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMsg(1, []byte(".......ping.......")))
		if err != nil {
			fmt.Println("failed to pack message", err)
		}
		conn.Write(msg)

		// read server message
		head := make([]byte, dp.HeadLen())
		if _, err = io.ReadFull(conn, head); err != nil {
			fmt.Println("failed to read server head", err)
			break
		}
		serverMsg, err := dp.UnPack(head)
		if err != nil {
			fmt.Println("failed to unpack server message head", err)
			break
		}
		if serverMsg.Len() > 0 {
			data := make([]byte, serverMsg.Len())
			if _, err = io.ReadFull(conn, data); err != nil {
				fmt.Println("failed to unpack server message data", err)
				break
			}
			fmt.Println("get server msg id: ", serverMsg.Id(), string(data))
		}
		time.Sleep(1 * time.Second)
	}
}
