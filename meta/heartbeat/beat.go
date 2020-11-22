package heartbeat

import (
	"fmt"
	"time"

	"shaos/util/redisC"
	"shaos/meta/conf"
)

func StartHearBeat() {
	for {
		err := redisC.Conn.String.Set(0, "meta:"+*conf.BaseConfig.ServerName, "localhost:"+*conf.BaseConfig.Port, 5).Error()
		if err != nil {
			fmt.Println("pub heartbeat err:", err)
		}
		time.Sleep(3 * time.Second)
	}
}
