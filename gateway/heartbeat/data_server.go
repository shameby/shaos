package heartbeat

import (
	"fmt"

	"math/rand"
	"shaos/util/redisC"
)

func getDataServers() []string {
	list, err := redisC.Conn.Key.Keys(0, "dh:*").Strings()
	if err != nil {
		fmt.Println("err when get data_server list")
		return nil
	}
	return list
}

func ChooseRandomDataServer() (appKey string, host string) {
	var (
		err  error
	)
	for rpTime := 0; rpTime < 3; rpTime++ {
		ds := getDataServers()
		n := len(ds)
		if n == 0 {
			continue
		}
		appKey = ds[rand.Intn(n)]
		host, err = redisC.Conn.String.Get(0, ds[rand.Intn(n)]).String()
		if err != nil {
			continue
		}
		if host != "" {
			break
		}
	}
	return
}
