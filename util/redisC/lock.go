package redisC

import (
	"fmt"

	"time"
)

const (
	DefaultTimeOut = 10
)

type Lock struct {
	timeout int64
	key     string
	conn    *RedigoPack
}

func (lock *Lock) Lock(token string) (ok bool, err error) {
	var (
		success bool
		ttlTime int64
		end     time.Time
	)
	end = time.Now().Add(time.Second * time.Duration(lock.timeout))
	for time.Now().Before(end) {
		if success, err = lock.conn.String.SetNX(7, lock.getKey(), token, lock.timeout).Bool(); success {
			return true, nil
		} else if ttlTime, _ = lock.conn.Key.TTL(7, lock.getKey()).Int64(); ttlTime == -1 {
			lock.conn.Key.Expire(7, lock.getKey(), lock.timeout)
		}
		time.Sleep(time.Microsecond)
	}

	return
}

func (lock *Lock) Unlock() (err error) {
	_, err = lock.conn.Key.Del(7, lock.getKey()).Result()
	return
}

func (lock *Lock) getKey() string {
	return fmt.Sprintf("redislock:%s", lock.key)
}

func NewLock(key string) (lock *Lock) {
	return NewLockWithTimeout(key, DefaultTimeOut)
}

func NewLockWithTimeout(key string, timeout int64) (lock *Lock) {
	lock = &Lock{timeout, key, Conn}

	return
}
