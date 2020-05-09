package znet

import (
	"fmt"
	"github.com/sunmeng90/zinx/utils"
	"github.com/sunmeng90/zinx/ziface"
	"net"
)

type Server struct {
	Name          string
	Version       string
	IpVersion     string
	IP            string
	Port          int
	MessageHandle ziface.IMessageHandle
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server name: %s, Version: %s, Ip: %s, Port: %d, MaxConn: %d, MaxPacketSize: %d\n",
		s.Name, s.Version, s.IP, s.Port, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPacketSize) // TODO make config consistent
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("Resolve tc address failed", err)
		return
	}

	tcp, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("Can not listen to address, error: ", err)
		return
	}
	fmt.Println("Server start successfully, listen on ", s.Port)
	var connId uint32 = 0
	for {
		accept, err := tcp.AcceptTCP()
		if err != nil {
			fmt.Println("Accept, error: ", err)
			continue
		}
		connId++
		go func() {
			// wrap a socket and business handler in a connection
			NewConn(accept, connId, s.MessageHandle).Start()
		}()
	}
}

func (s *Server) Stop() {
	panic("implement me")
}

func (s *Server) Serve() {
	go func() {
		s.Start()
	}()

	// TODO
	// do something
	select {} // blocking
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	fmt.Println("Add a new router")
	s.MessageHandle.AddRouter(msgId, router)
}

func NewServer(name string) *Server {
	return &Server{
		Name:          utils.GlobalObject.Name,
		IpVersion:     "tcp4",
		Version:       utils.GlobalObject.Version,
		IP:            utils.GlobalObject.Host,
		Port:          utils.GlobalObject.TcpPort,
		MessageHandle: NewMessageHandle(),
	}
}
