package redis

import "time"

func (c *client) Ping() error {
	_, err := c.redisClient.Ping().Result()
	return err
}

func (c *client) Set(key string, value interface{}, expiration time.Duration) error {
	_, err := c.redisClient.Set(key, value, expiration).Result()
	return err
}

func (c *client) Get(key string) (interface{}, error) {
	result, err := c.redisClient.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}
