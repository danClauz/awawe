package redis

import (
	"encoding/json"
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
	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if _, err := c.redisClient.Set(key, byteValue, expiration).Result(); err != nil {
		return err
	}

	return nil
}

func (c *sdk) Get(key string) (interface{}, error) {
	result, err := c.redisClient.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
