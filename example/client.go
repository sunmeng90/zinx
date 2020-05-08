package main

import (
	"fmt"
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
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("failed to write data to server", err)
			return
		}

		buf := make([]byte, 512)
		nRead, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read data from server")
			return
		}
		fmt.Println("Got server Data: ", string(buf[:nRead]))
		time.Sleep(1 * time.Second)
	}
}
