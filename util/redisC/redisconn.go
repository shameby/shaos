package redisC

import (
	"time"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	pool        *redis.Pool
	redisHost         = "118.24.255.170:6379"
	redisPass         = "hansliaoEva613"
	redisPrefix int64 = 20
)

//newRedisPool:创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   10,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				panic(fmt.Sprintf("redis init err: %v", err))
			}
			if _, err := c.Do("AUTH", redisPass); err != nil {
				c.Close()
				panic(fmt.Sprintf("redis init err: %v", err))
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

//初始化redis连接池
func init() {
	pool = newRedisPool()
}
