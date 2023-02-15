package config

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	ctx     = context.Background()
	DBRedis *redis.Client
)

func OpenRedisPool() {
	DBRedis = ConnectRedis()
}

func ConnectRedis() *redis.Client {
	opt := &redis.Options{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
		Password:     "",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	}

	rdb := redis.NewClient(opt)
	_, err := rdb.Ping().Result()

	if err != nil {
		panic(err)
	}

	return rdb
}
