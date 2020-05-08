package znet

import (
	"fmt"
	"net"
)

type Server struct {
	Name      string
	Version   string
	IpVersion string
	IP        string
	Port      int
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

	for {
		accept, err := tcp.AcceptTCP()
		if err != nil {
			fmt.Println("Accept, error: ", err)
			continue
		}

		go func() {
			for {
				buf := make([]byte, 512)
				nbytes, err := accept.Read(buf)
				if err != nil {
					fmt.Println("Read data on socket, error: ", err)
					continue // stop current conn
				}
				if _, err := accept.Write(buf[0:nbytes]); err != nil {
					fmt.Println("Failed to write back, error: ", err)
					continue
				}
			}
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

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IpVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
