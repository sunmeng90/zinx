package ziface

type IDataPack interface {
	HeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	UnPack(packet []byte) (IMessage, error)
}
