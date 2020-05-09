package znet

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"net"
)

type Server struct {
	Name      string
	Version   string
	IpVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
}

func (s *Server) Start() {
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
			NewConn(accept, connId, s.Router).Start()
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

func (s *Server) AddRouter(router ziface.IRouter) {
	fmt.Println("Add a new router")
	s.Router = router
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
}
