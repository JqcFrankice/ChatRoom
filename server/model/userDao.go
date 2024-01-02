package model

import (
	"chatroom/server/db"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// InitUserDao 函数用于初始化用户数据访问对象
func InitUserDao() {
	MyUserDao = NewUserDao(db.Pool)
}

// MyUserDao 全局变量，表示用户数据访问对象
var (
	MyUserDao *UserDao
)

// UserDao 结构体定义了用户数据访问对象
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 函数用于创建一个新的用户数据访问对象
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// getUserById 方法用于通过用户ID从 Redis 中获取用户信息
func (this *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "user", userId))
	if err != nil {
		// 处理用户不存在的情况
		if err == redis.ErrNil {
			err = ERROR_USER_NOTXEISTS
		}
		return
	}
	user = &User{}
	// 将获取到的 JSON 数据反序列化为 User 结构体
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Marshal([]byte(res) srr = ", err)
		return
	}
	return
}

// Login 方法用于处理用户登录逻辑
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	// 通过用户ID获取用户信息
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 校验用户密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

// Register 方法用于处理用户注册逻辑
func (this *UserDao) Register(userId int, userPwd string, userName string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	// 判断用户是否已经存在
	_, err = this.getUserById(conn, userId)
	if err == nil {
		// 用户已存在的错误处理
		err = ERROR_USER_EXISTS
		return
	}

	// 创建新用户并设置属性值
	user.UserId = userId
	user.UserPwd = userPwd
	user.UserName = userName

	// 将用户信息序列化为 JSON 数据
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Scanf("json.Marshal(user) err = ", err)
	}

	fmt.Printf("json.Marshal(user) %v\n", string(data))

	// 将用户信息保存到 Redis 中
	_, err = conn.Do("Hset", "user", userId, string(data))
	if err != nil {
		fmt.Println("报错报错")
		return
	}
	return
}
