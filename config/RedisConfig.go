package config

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var rdb *redis.Client
var Once sync.Once //只会运行一次

func BuildRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
