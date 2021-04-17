package redis

import (
	"getaway/config"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func RsInit() (err error) {
	cfg := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
		DB:   0,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
