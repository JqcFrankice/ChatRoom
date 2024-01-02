package model

// User 结构体定义了用户对象的属性
type User struct {
	UserId   int    `json:"user_id"`   // 用户ID
	UserPwd  string `json:"user_pwd"`  // 用户密码
	UserName string `json:"user_name"` // 用户名
}
