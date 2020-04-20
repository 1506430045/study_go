package util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"study_go/app/lib/db/redis/kv"
	"time"
)

type RedisLimiter interface {
	Allow() error
	AllowN(n int) error
}

type RedisLimit struct {
	key   string
	limit int
	d     time.Duration
}

const RedisLimitKeyPrefix = "redis:limiter"

func NewRedisLimiter(key string, limit int, d time.Duration) (*RedisLimit, error) {
	limiter := &RedisLimit{
		key:   key,
		limit: limit,
		d:     d,
	}
	_, err := redis.Int(kv.Do("SETNX", limiter.getKey(), limiter.limit))
	if err != nil {
		return nil, err
	}
	ticker := time.NewTicker(d) //定时恢复令牌数量
	go func() {
		for t := range ticker.C {
			_, err = redis.Int(kv.Do("SET", limiter.getKey(), limit))
			fmt.Println("tick_at", t, "err", err)
		}
	}()
	return limiter, nil
}

func (r *RedisLimit) getKey() string {
	return fmt.Sprintf("%s:%s", RedisLimitKeyPrefix, r.key)
}

func (r *RedisLimit) Allow() error {
	num, err := redis.Int(kv.Do("DECR", r.getKey()))
	if err != nil {
		return err
	}
	if num < 0 {
		return fmt.Errorf("limit")
	}
	return nil
}

//借助lua脚本
func (r *RedisLimit) AllowN(n int) error {
	c := kv.ConnPool().Get()
	defer c.Close()
	return nil
}
