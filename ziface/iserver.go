package ziface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgId uint32, router IRouter)
	ConnManager() IConnManager
	SetOnConnStart(onConnStart func(conn IConn))
	SetOnConnStop(onConnStop func(conn IConn))
	CallOnConnStart(conn IConn)
	CallOnConnStop(conn IConn)
}
