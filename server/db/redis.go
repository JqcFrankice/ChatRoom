package db

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// Pool 全局变量，用于存储 Redis 连接池
var Pool *redis.Pool

// InitPool 初始化 Redis 连接池
func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	// 创建 Redis 连接池对象
	Pool = &redis.Pool{
		// Dial 函数用于创建和配置一个新的连接
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
		// MaxIdle 是连接池中最大空闲连接数
		MaxIdle: maxIdle,
		// MaxActive 是连接池中最大活动连接数
		MaxActive: maxActive,
		// IdleTimeout 是连接池中连接的最大空闲时间
		IdleTimeout: idleTimeout,
	}
}
