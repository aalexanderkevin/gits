package storage

import (
	"fmt"
	"test3/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewRedis(redisConfig config.RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: time.Duration(redisConfig.IdleTimeout) * time.Second,
		MaxActive:   redisConfig.MaxActive,
		Dial: func() (redis.Conn, error) {
			network := "tcp"
			address := fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)
			conn, err := redis.Dial(network, address)

			// Connection error handling
			if err != nil {
				fmt.Printf("Failed to initializing the redis pool, redis.Dial, err:%s, network:%s, address:%s"+err.Error(), network, address)
				return nil, err
			}
			return conn, err
		},
	}
}
