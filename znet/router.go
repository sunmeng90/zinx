package znet

import (
	"github.com/sunmeng90/zinx/ziface"
)

type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(req ziface.IRequest) {
}

func (b *BaseRouter) Handle(req ziface.IRequest) {
}

func (b *BaseRouter) PostHandle(req ziface.IRequest) {
}
