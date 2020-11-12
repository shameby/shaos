package heartbeat

import (
	"fmt"
	"time"

	"shaos/util/redisC"
)

func StartHearBeat() {
	for {
		_, err := redisC.Conn.Ps.Publish(0, "heartbeat", "localhost:8081").Int()
		if err != nil {
			fmt.Println("pub heartbeat err:", err)
		}
		time.Sleep(5 * time.Second)
	}
}
