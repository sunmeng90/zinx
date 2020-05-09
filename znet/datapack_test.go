package znet

import (
	"fmt"
	"github.com/sunmeng90/zinx/ziface"
	"io"
	"net"
	"reflect"
	"testing"
)

func TestDataPack_2(t *testing.T) {
	serv, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("failed to listen to 7777", err)
		return
	}

	go func() {
		for {
			conn, err := serv.Accept()
			if err != nil {
				fmt.Println("failed to get connection", err)
				continue
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.HeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("failed to read  header", err)
						break
					}

					msg, err := dp.UnPack(headData) // unpack head
					if err != nil {
						fmt.Println("failed to unpack header")
						return
					}
					if msg.Len() > 0 {
						msgData := make([]byte, msg.Len())
						_, err := io.ReadFull(conn, msgData)
						if err != nil {
							fmt.Println("failed to unpack data")
							return
						}
						msg.SetData(msgData) // TODO: inconsistency, setter is unnecessary
					}

					fmt.Println("receive msg ", msg.Id(), " data: ", string(msg.Data()))

				}

			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("failed to connect to server", err)
	}
	dp := NewDataPack()
	msg1 := &Message{
		id:   1,
		len:  5,
		data: []byte("hello"),
	}
	msg2 := &Message{
		id:   2,
		len:  5,
		data: []byte("world"),
	}
	pack1Bytes, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}
	pack2Bytes, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error", err)
	}

	conn.Write(append(pack1Bytes, pack2Bytes...))
	select {}
}

func TestDataPack_HeadLen(t *testing.T) {
	tests := []struct {
		name string
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataPack{}
			if got := d.HeadLen(); got != tt.want {
				t.Errorf("HeadLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataPack_Pack(t *testing.T) {
	type args struct {
		msg ziface.IMessage
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataPack{}
			got, err := d.Pack(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataPack_UnPack(t *testing.T) {
	type args struct {
		packet []byte
	}
	tests := []struct {
		name    string
		args    args
		want    ziface.IMessage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DataPack{}
			got, err := d.UnPack(tt.args.packet)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnPack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnPack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDataPack(t *testing.T) {
	tests := []struct {
		name string
		want *DataPack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataPack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataPack() = %v, want %v", got, tt.want)
			}
		})
	}
}
