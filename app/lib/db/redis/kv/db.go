package kv

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var __ConnPool *redis.Pool

const RedisHost = "127.0.0.1"
const RedisPort = 6379
const RedisPassword = ""

func init() {
	defer func() {
		if err := recover(); err != nil {
			_ = Close()
			panic(err)
		}

		go monitorLoop(100)
	}()

	__ConnPool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", RedisHost, RedisPort),
				redis.DialPassword(RedisPassword),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(2*time.Second),
				redis.DialReadTimeout(2*time.Second),
				redis.DialWriteTimeout(2*time.Second),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: time.Second,
		Wait:        true,
	}

	if _, err := Do("PING"); err != nil {
		panic(err)
	}
}
