package znet

type Message struct {
	id   uint32
	len  uint32
	data []byte
}

func NewMsg(id uint32, data []byte) *Message {
	return &Message{
		id:   id,
		len:  uint32(len(data)),
		data: data,
	}
}

func (m *Message) Id() uint32 {
	return m.id
}

func (m *Message) Len() uint32 {
	return m.len
}

func (m *Message) Data() []byte {
	return m.data
}

func (m *Message) SetId(id uint32) {
	m.id = id
}

func (m *Message) SetLen(len uint32) {
	m.len = len
}

func (m *Message) SetData(data []byte) {
	m.data = data
}
