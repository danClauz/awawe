package redis

import (
	config "awawe/configuration"
	"github.com/go-redis/redis"
	"time"
)

type client struct {
	redisClient *redis.Client
}

type Client interface {
	Ping() error
	Set(key string, value interface{}, expiration time.Duration) error
}

func NewRedisClient() Client {
	redisConfig := config.GetRedisConfig()
	return &client{
		redisClient: redis.NewClient(&redis.Options{
			Addr:         redisConfig.Address,
			Password:     redisConfig.Password,
			DB:           redisConfig.DB,
			ReadTimeout:  time.Duration(redisConfig.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(redisConfig.WriteTimeout) * time.Second,
			PoolSize:     redisConfig.PoolSize,
		}),
	}
}
