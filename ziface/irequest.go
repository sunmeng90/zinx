package ziface

type IRequest interface {
	// TODO reorganize move interface and struct into once package to avoid import cycle
	Conn() IConn // can't use znet.Conn, otherwise will cause import cycle
	Data() []byte
	MsgId() uint32
}
