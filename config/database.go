package config

import (
	"github.com/redis/go-redis/v9"
)

func connectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     urlRedis,
		Password: "",
		DB:       0,
	})
}
