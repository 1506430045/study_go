package kv

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

func monitorLoop(maxActive int) {
	ticker := time.NewTicker(time.Minute * 5)
	defer ticker.Stop() //目前永远结束不了, 后期优化这个吧.

	monitor(maxActive)
	for range ticker.C {
		monitor(maxActive)
	}
}

func monitor(maxActive int) {
	pool := ConnPool()
	active := pool.ActiveCount()
	idle := pool.IdleCount()
	inUse := active - idle

	if 5*active > 4*maxActive { //当前连接数大于总连接数的80%
		log.Warn(fmt.Sprintf("func:redis-monitor:ActiveCount:%d:IdleCount:%d:InUseCount:%d", active, idle, inUse))
	} else {
		log.Info(fmt.Sprintf("func:redis-monitor:ActiveCount:%d:IdleCount:%d:InUseCount:%d", active, idle, inUse))
	}

	if _, err := Do("PING"); err != nil {
		log.Error("Redis.Ping", err.Error())
	}
}
