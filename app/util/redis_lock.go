package util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"math/rand"
	"study_go/app/lib/db/redis/kv"
	"time"
)

//redis锁实现
type RedisLock interface {
	Lock() (string, error)
	Unlock() error
}

type RedLock struct {
	LockKey string
	Expire  int64
}

//添加锁
func (r *RedLock) Lock() (randomStr string, err error) {
	randomStr = Md5(fmt.Sprintf("random:str:%d:%d", time.Now().Unix(), rand.Int())) //只是示例
	ok, err := redis.String(kv.Do("SET", r.LockKey, randomStr, "EX", r.Expire, "NX"))
	if err != nil || ok != "OK" {
		return "", errors.Wrap(err, "get lock failed")
	}
	return
}

//释放锁
func (r *RedLock) Unlock(randomStr string) error {
	c := kv.ConnPool().Get()
	defer c.Close()
	luaScript := `if redis.call("get",KEYS[1]) == ARGV[1] then
						return redis.call("del",KEYS[1])
                  else
						return 
                  end`
	goScript := redis.NewScript(1, luaScript)
	re, err := redis.Int(goScript.Do(c, r.LockKey, randomStr))
	if err != nil || re != 1 {
		return errors.Wrap(err, "unlock failed")
	}
	return nil
}
