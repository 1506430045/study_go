package util

import (
	"github.com/garyburd/redigo/redis"
	"study_go/app/lib/db/redis/kv"
)

type RedisBiter interface {
	Get(key int64) int
	Set(key int64) bool
}

type RedisBit struct {
	key       string
	SetSuffix func() string
	SetExpire func() int
}

//初始化redis bitmap
func NewRedisBit(key string, SuffixKey func() string, Expire func() int) (*RedisBit, error) {
	r := &RedisBit{
		key:       key,       //key
		SetSuffix: SuffixKey, //后缀
		SetExpire: Expire,    //到期
	}

	exists, err := redis.Int(kv.Do("EXISTS", r.getKey()))
	if err != nil {
		return r, err
	}
	if exists == 0 { //如果不存在 则初始化
		_, err := redis.Int(kv.Do("SETBIT", r.getKey(), 1, 0))
		if err != nil {
			return r, err
		}
		_, err = redis.Int(kv.Do("EXPIRE", r.getKey(), Expire()))
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

func (r *RedisBit) getKey() string {
	return r.key + r.SetSuffix()
}

func (b *RedisBit) Get(key int64) int {
	val, err := redis.Int(kv.Do("GETBIT", b.getKey(), key))
	if err != nil {
		return 0
	}
	return val
}

func (b *RedisBit) Set(key int64) bool {
	val, err := redis.Int(kv.Do("SETBIT", b.getKey(), key, 1))
	if err != nil {
		return false
	}
	return val == 0
}
