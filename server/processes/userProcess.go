package processes

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// UserProcess 结构体定义了用户处理过程
type UserProcess struct {
	Conn   net.Conn // 连接对象
	UserId int      // 用户ID
}

// ServerProcessLogin 方法用于处理用户登录
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 解析客户端发送的登录消息
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Date), &loginmes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
	}

	// 构建响应消息
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	// 调用模型层处理用户登录逻辑
	user, err := model.MyUserDao.Login(loginmes.UserId, loginmes.UserPwd)
	if err != nil {
		// 处理登录失败情况
		if err == model.ERROR_USER_EXISTS {
			loginResMes.Code = 300
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 400
		} else if err == model.ERROR_USER_NOTXEISTS {
			loginResMes.Code = 500
		}
		loginResMes.Err = err.Error()
	} else {
		// 处理登录成功情况
		loginResMes.Code = 200
		this.UserId = loginmes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginmes.UserId)
		for id, _ := range userMgr.OnlineUser {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Printf("用户消息user%v", user)
	}

	// 将响应消息发送给客户端
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal faile err=", err)
		return
	}
	resMes.Date = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	tr := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("服务器发送数据失败")
	}

	return
}

// ServerProcessRegister 方法用于处理用户注册
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 解析客户端发送的注册消息
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Date), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshal err = ", err)
	}

	// 构建响应消息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// 调用模型层处理用户注册逻辑
	user, err := model.MyUserDao.Register(registerMes.UserId, registerMes.UserPwd, registerMes.UserName)
	if err != nil {
		// 处理注册失败情况
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 300
		} else if err == model.ERROR_USER_PWD {
			registerResMes.Code = 400
		} else if err == model.ERROR_USER_NOTXEISTS {
			registerResMes.Code = 500
		}
		registerResMes.Err = err.Error()
	} else {
		// 处理注册成功情况
		registerResMes.Code = 200
		fmt.Printf("用户消息user%v", user)
	}

	// 将响应消息发送给客户端
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.marshal faile err=", err)
		return
	}
	resMes.Date = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	tr := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("服务器发送数据失败")
	}

	return
}

// NotifyOthersOnlineUser 方法用于通知其他用户上线信息
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 打印提示信息
	fmt.Println("通知其他在线用户上线")

	// 遍历在线用户列表
	for id, up := range userMgr.OnlineUser {
		// 跳过自己
		if id == userId {
			continue
		}
		// 向其他用户发送上线通知
		up.NotityMetoOther(userId)
	}
}

// NotityMetoOther 方法用于通知其他用户有人上线
func (this *UserProcess) NotityMetoOther(userId int) {
	// 创建消息结构体
	var mes message.Message
	mes.Type = message.NotyfyUserStatusMesType

	// 创建用户状态变更消息结构体
	var notifyUserStatues message.NotyfyUserStatusMes
	notifyUserStatues.UserId = userId
	notifyUserStatues.Status = message.UserOnline

	// 将用户状态变更消息结构体序列化为 JSON 数据
	data, err := json.Marshal(notifyUserStatues)
	if err != nil {
		fmt.Println("data, err := json.Marshal(notifyUserStatues) err=", err)
		return
	}

	// 设置消息内容
	mes.Date = string(data)

	// 将消息结构体序列化为 JSON 数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("data, err := json.Marshal(notifyUserStatues) err=", err)
		return
	}

	// 创建 Transfer 对象，用于向其他用户发送消息
	tr := &utils.Transfer{
		Conn: this.Conn,
	}

	// 向其他用户发送消息
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("err = tr.WritePkg(data) err=", err)
		return
	}
}
