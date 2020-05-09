package znet

import (
	"fmt"
	"github.com/sunmeng90/zinx/utils"
	"github.com/sunmeng90/zinx/ziface"
	"strconv"
)

type MessageHandle struct {
	Apis           map[uint32]ziface.IRouter
	TaskQueues     []chan ziface.IRequest
	WorkerPoolSize uint32
}

func NewMessageHandle() *MessageHandle {
	return &MessageHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueues:     make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

func (m *MessageHandle) StartWorkerPool() {
	fmt.Println("Start worker pool")
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueues[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxTaskSize)
		go m.startWorker(i, m.TaskQueues[i])
	}
}

func (m *MessageHandle) startWorker(i int, taskQueue chan ziface.IRequest) {
	fmt.Println("Start worker ", i)
	for {
		select {
		case req := <-taskQueue:
			fmt.Println("worker ", i, "take task from queue")
			m.DoMsgHandler(req)
		}
	}
}

// connection(request+) -> queue
func (m *MessageHandle) SendReqToQueue(req ziface.IRequest) {
	queueIdx := req.Conn().GetConnID() % m.WorkerPoolSize
	fmt.Println("send request in connection ", req.Conn().GetConnID(), " to queue ", queueIdx)
	m.TaskQueues[queueIdx] <- req
}
