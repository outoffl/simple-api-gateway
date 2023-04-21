package cache

import (
	"api-gateway/internal/cfg"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var redisPool *redis.Pool

func InitPoolRedis() {
	rdsCfg := cfg.LoadRedisCfg()
	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", rdsCfg.Host, rdsCfg.Port))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			return err
		},
		MaxIdle:         1,
		MaxActive:       4,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
func CheckPool() {
	fmt.Println(redisPool.ActiveCount())
	fmt.Println(redisPool.IdleCount())
}

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	pool := redisPool.Get()
	defer func() {
		fmt.Println("closeConn here")
		if err := pool.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	return pool.Do(commandName, args...)
}
