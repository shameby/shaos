package redisC

import (
	"time"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type psRds struct {
}

// Publish 将信息发送到指定的频道，返回接收到信息的订阅者数量
func (p *psRds) Publish(db int64, channel, message string) *Reply {
	c := pool.Get()
	defer c.Close()
	return getReply(c.Do("PUBLISH", channel, message))
}

// Subscribe 订阅给定的一个或多个频道的信息。
// 支持redis服务停止或网络异常等情况时，自动重新订阅。
// 一般的程序都是启动后开启一些固定channel的订阅，也不会动态的取消订阅，这种场景下可以使用本方法。
// 复杂场景的使用可以直接参考 https://godoc.org/github.com/gomodule/redigo/redis#hdr-Publish_and_Subscribe
func (p *psRds) Subscribe(onMessage func(channel string, data []byte), channels ...string) error {
	c := pool.Get()
	psc := redis.PubSubConn{Conn: c}
	err := psc.Subscribe(redis.Args{}.AddFlat(channels)...)
	// 如果订阅失败，休息1秒后重新订阅（比如当redis服务停止服务或网络异常）
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second)
		return p.Subscribe(onMessage, channels...)
	}
	quit := make(chan int, 1)

	// 处理消息
	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				if len(v.Data) != 0 {
					go onMessage(v.Channel, v.Data)
				}
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				quit <- 1
				fmt.Println(v.Error())
				return
			}
		}
	}()

	// 异常情况下自动重新订阅
	go func() {
		<-quit
		time.Sleep(time.Second)
		psc.Close()
		c.Close()
		p.Subscribe(onMessage, channels...)
	}()
	return err
}
