package kv

import "github.com/garyburd/redigo/redis"

//ConnPool 返回 redis.Pool.
//除非必要一般不建议用这个函数, 用本库封装好的函数操作数据库.
func ConnPool() *redis.Pool {
	return __ConnPool
}

//Close释放连接资源.
func Close() error {
	if __ConnPool != nil {
		return __ConnPool.Close()
	}
	return nil
}

//Do执行redis命令
//NOTE除非有必要(比如在一个函数内部需要执行多次redis操作), 否则请用该函数执行所有的操作, 这样能有效避免忘记释放资源.
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := __ConnPool.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}
