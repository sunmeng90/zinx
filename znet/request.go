package znet

import "github.com/sunmeng90/zinx/ziface"

type Request struct {
	conn ziface.IConn // use pointer or value
	data []byte
}

func (r *Request) Conn() ziface.IConn {
	return r.conn
}
func (r *Request) Data() []byte {
	return r.data
}
