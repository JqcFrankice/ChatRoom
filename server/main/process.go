package main

import (
	"chatroom/common/message"
	"chatroom/server/processes"
	"chatroom/server/utils"
	"fmt"
	"net"
)

// Process 结构体定义了处理客户端连接的对象
type Process struct {
	Conn net.Conn // 与客户端建立的连接对象
}

// ServerProcessMes 方法用于处理客户端发送的不同类型消息
func (this *Process) ServerProcessMes(mes *message.Message) (err error) {
	// 根据消息类型进行不同的处理
	switch mes.Type {
	case message.LoginMesType:
		// 如果是登录消息，则创建用户处理对象，调用登录处理函数
		us := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = us.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 如果是注册消息，则创建用户处理对象，调用注册处理函数
		us := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = us.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 如果是群发消息，则创建群发消息处理对象，调用群发消息处理函数
		fmt.Println("群发信息")
		smp := &processes.SmsProcess{}
		smp.SendGroupMes(mes)
	default:
		// 其他消息类型，打印错误信息
		fmt.Println("消息类型不存在 无法处理")
	}
	return
}

// processAll 方法用于循环处理客户端发送的消息
func (this *Process) processAll() {
	for {
		// 创建 Transfer 对象，用于接收客户端发送的消息
		ut := &utils.Transfer{
			Conn: this.Conn,
		}
		// 读取客户端发送的消息
		mes, err := ut.ReadPkg()
		if err != nil {
			fmt.Println("接受buf err=", err)
			return
		}
		fmt.Println("mes=", mes)
		// 处理客户端发送的消息
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("返回消息失败")
		}
	}
}
