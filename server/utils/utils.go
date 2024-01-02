package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// Transfer 结构体定义了数据传输的结构
type Transfer struct {
	Conn net.Conn   // 连接对象
	Buf  [8090]byte // 数据缓冲区
}

// WritePkg 方法用于发送消息数据包
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))

	// 将消息长度写入缓冲区的前4个字节，使用大端序
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	// 发送消息长度给对方
	_, err = this.Conn.Write(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write() err=", err)
		return
	}

	// 发送实际消息数据给对方
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write() err=", err)
		return
	}

	return
}

// ReadPkg 方法用于接收消息数据包
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	// 先读取前4个字节，获取消息的长度
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("接收buf err=", err)
		return
	}
	fmt.Println("接收buf =", this.Buf[:4])

	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(this.Buf[:4])

	// 根据消息长度读取实际的消息数据
	n, err := this.Conn.Read(this.Buf[:dataLen])
	if n != int(dataLen) || err != nil {
		fmt.Println("接收buf err=", err)
		return
	}

	// 将接收到的消息数据反序列化为 Message 结构体
	err = json.Unmarshal(this.Buf[:dataLen], &mes)
	if err != nil {
		fmt.Println("接收buf err=", err)
		return
	}

	return
}
