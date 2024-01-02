package processes

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// SmsProcess 结构体定义了消息处理过程
type SmsProcess struct{}

// SendGroupMes 方法用于发送群聊消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 解析客户端发送的群聊消息
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Date), &smsMes)
	if err != nil {
		fmt.Println("err = json.Unmarshal([]byte(mes.Date), smsMes) err=", err)
		return
	}

	// 将消息序列化为 JSON 数据
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("err = json.Unmarshal([]byte(mes.Date), smsMes) err=", err)
		return
	}

	// 遍历在线用户，向除发送者外的其他用户发送消息
	for id, up := range userMgr.OnlineUser {
		if id == smsMes.User.UserId {
			continue
		}
		this.SendMesToOther(data, up.Conn)
	}
	return
}

// SendMesToOther 方法用于向其他用户发送消息
func (this *SmsProcess) SendMesToOther(data []byte, conn net.Conn) {
	// 使用 Transfer 结构体进行数据发送
	tr := &utils.Transfer{
		Conn: conn,
	}

	// 将消息发送给指定用户
	err := tr.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToOther err=", err)
	}
	fmt.Println("群发消息结束")
}
