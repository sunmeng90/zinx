package ziface

type IMessageHandle interface {
	// dispatch request
	DoMsgHandler(req IRequest)
	AddRouter(msgId uint32, router IRouter)
}
