package processes

import "fmt"

var (
	userMgr *UserMgr
)

// UserMgr 结构体定义了用户管理器
type UserMgr struct {
	OnlineUser map[int]*UserProcess // 在线用户列表，键为用户ID，值为用户处理过程对象
}

// init 函数在包被导入时执行，用于初始化全局变量 userMgr
func init() {
	userMgr = &UserMgr{
		OnlineUser: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUser 方法用于添加在线用户到用户管理器
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.OnlineUser[up.UserId] = up
}

// DelOnlineUser 方法用于从用户管理器删除在线用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.OnlineUser, userId)
}

// GetAllOnlineUser 方法用于获取所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.OnlineUser
}

// GetOnlineUserById 方法用于通过用户ID获取在线用户处理过程对象
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.OnlineUser[userId]
	if !ok {
		// 用户不存在的错误处理
		err = fmt.Errorf("用户不存在 %d \n", userId)
		return
	}
	return
}
