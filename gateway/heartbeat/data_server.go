package heartbeat

import (
	"time"
	"sync"
	"math/rand"
)

type dst struct {
	m map[string]time.Time
	sync.RWMutex
}

var dataServers *dst

func init() {
	dataServers = &dst{
		m: make(map[string]time.Time),
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		dataServers.Lock()
		for s, t := range dataServers.m {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers.m, s)
			}
		}
		dataServers.Unlock()
	}
}

func getDataServers() []string {
	dataServers.RLock()
	defer dataServers.RUnlock()
	ds := make([]string, 0, len(dataServers.m))
	for s := range dataServers.m {
		ds = append(ds, s)
	}
	return ds
}

func ChooseRandomDataServer() string {
	ds := getDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
