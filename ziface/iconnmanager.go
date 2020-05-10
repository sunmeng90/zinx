package ziface

type IConnManager interface {
	Add(conn IConn)
	Remove(conn IConn)
	Get(id uint32) (IConn, error)
	Len() int
	Clear()
}
