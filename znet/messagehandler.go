package znet

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"strconv"
)

type MessageHandle struct {
	Apis map[uint32]ziface.IRouter
}

func NewMessageHandle() *MessageHandle {
	return &MessageHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (m *MessageHandle) DoMsgHandler(req ziface.IRequest) {
	router, ok := m.Apis[req.MsgId()]
	if !ok {
		fmt.Println("router not found for message id:", req.MsgId())
		return
	}
	router.PreHandle(req)
	router.Handle(req)
	router.PostHandle(req)
}

func (m *MessageHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		panic("repeat api call, message id: " + strconv.Itoa(int(msgId)))
	}
	m.Apis[msgId] = router
	fmt.Println("add new router for msg: ", msgId)
}
