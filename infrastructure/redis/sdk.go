package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type sdk struct {
	redisClient *redis.Client
}

func (c *sdk) Ping() error {
	_, err := c.redisClient.Ping().Result()
	return err
}

func (c *sdk) Set(key string, value interface{}, expiration time.Duration) error {
	_, err := c.redisClient.Set(key, value.(string), expiration).Result()
	return err
}

func (c *sdk) Get(key string) (interface{}, error) {
	result, err := c.redisClient.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
