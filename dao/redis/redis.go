package redis

import (
	"getaway/config"
	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
	Nil         = redis.Nil
)

func RsInit() (err error) {
	cfg := config.GetConfig()
	RedisClient = redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
		DB:   0,
	})
	_, err = RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = RedisClient.Close()
}
