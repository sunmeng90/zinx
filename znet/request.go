package znet

import "github.com/sunmeng90/zinx/ziface"

type Request struct {
	conn ziface.IConn // use pointer or value
	msg  ziface.IMessage
}

func (r *Request) Conn() ziface.IConn {
	return r.conn
}
func (r *Request) Data() []byte {
	return r.msg.Data()
}

func (r *Request) MsgId() uint32 {
	return r.msg.Id()
}
