package model

import "errors"

// 自定义错误常量，表示用户不存在
var (
	ERROR_USER_NOTXEISTS = errors.New("用户不存在。。。")
	// 自定义错误常量，表示用户已经存在
	ERROR_USER_EXISTS = errors.New("用户已经存在。。。")
	// 自定义错误常量，表示密码不正确
	ERROR_USER_PWD = errors.New("密码不正确")
)
