package ziface

type IMessage interface {
	Id() uint32
	Len() uint32
	Data() []byte
	SetId(id uint32)
	SetLen(len uint32)
	SetData(data []byte)
}
