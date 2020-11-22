package heartbeat

import (
	"fmt"
	"time"

	"shaos/util/redisC"
	"shaos/data/conf"
)

func StartHearBeat() {
	for {
		err := redisC.Conn.String.Set(0, "dh:"+conf.BaseConfig.AppKey, "localhost:"+*conf.BaseConfig.Port, 5).Error()
		if err != nil {
			fmt.Println("pub heartbeat err:", err)
		}
		time.Sleep(3 * time.Second)
	}
}
