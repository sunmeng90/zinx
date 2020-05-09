package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"io/ioutil"
)

type GlobalObj struct {
	TcpServer      ziface.IServer
	Host           string
	TcpPort        int
	Name           string
	Version        string
	MaxConn        int
	MaxPacketSize  uint32
	WorkerPoolSize uint32
	MaxTaskSize    uint32 // not the total size
}

var GlobalObject *GlobalObj

func init() {
	// default configuration
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0", //127.0.0.1 not working when there are multiple network interface card (NIC)
		TcpPort:        8999,
		Name:           "ZinxServerApp",
		Version:        "v0.5",
		MaxConn:        1000,
		MaxPacketSize:  4096,
		WorkerPoolSize: 10,
		MaxTaskSize:    1024,
	}
	// load custom config
	GlobalObject.reload()
}

func (g *GlobalObj) reload() {
	conf, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("failed to load custom configuration")
		panic(err)
	}

	err = json.Unmarshal(conf, &g)
	if err != nil {
		fmt.Println("failed to parse configuration")
		panic(err)
	}
}
