package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/sunmeng90/zinx/utils"
	"github.com/sunmeng90/zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) HeadLen() uint32 {
	// len(4 bytes) + id(4 bytes)
	return 8 // protocol
}

// len|id|data
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	// len
	if err := binary.Write(buf, binary.LittleEndian, msg.Len()); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Id()); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Data()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *DataPack) UnPack(packet []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(packet)

	msg := &Message{}
	if err := binary.Read(reader, binary.LittleEndian, &msg.len); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPacketSize > 0 && msg.Len() > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("exceeds max packet size")
	}

	if err := binary.Read(reader, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}

	return msg, nil
}
