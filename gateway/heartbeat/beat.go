package heartbeat

import (
	"time"

	"shaos/util/redisC"
)

func ListenHeartbeat() {
	redisC.Conn.Ps.Subscribe(handleMsg, "heartbeat")
	go removeExpiredDataServer()
}

func handleMsg(channel string, data []byte) {
	dataServerName := string(data)
	dataServers.Lock()
	defer dataServers.Unlock()
	dataServers.m[dataServerName] = time.Now()
}
