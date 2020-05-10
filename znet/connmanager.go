package znet

import (
	"errors"
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"strconv"
	"sync"
)

type ConnManager struct {
	connMap map[uint32]ziface.IConn
	lock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connMap: make(map[uint32]ziface.IConn),
		lock:    sync.RWMutex{},
	}
}

func (c *ConnManager) Add(conn ziface.IConn) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.connMap[conn.GetConnID()] = conn
	fmt.Println("add connection id", conn.GetConnID(), " to manager")
}

func (c *ConnManager) Remove(conn ziface.IConn) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.connMap, conn.GetConnID())
	fmt.Println("remove connection id", conn.GetConnID(), " from manager")
}

func (c *ConnManager) Get(id uint32) (ziface.IConn, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	conn, ok := c.connMap[id]
	if !ok {
		return nil, errors.New("connection " + strconv.Itoa(int(id)) + " not found")
	}
	return conn, nil
}

func (c *ConnManager) Len() int {
	return len(c.connMap)
}

func (c *ConnManager) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	for _, conn := range c.connMap {
		conn.Stop()
		delete(c.connMap, conn.GetConnID()) // TODO: is it safe for remove in iteration?
	}
	fmt.Println("clear all connections from manager")
}
