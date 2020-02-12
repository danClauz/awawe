package redis

import "time"

type SDK interface {
	Ping() error
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
}
