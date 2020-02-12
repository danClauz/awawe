package redis

import (
	iRedis "awawe/usecase/infrastructure/redis"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type client struct {
	Address      string
	Password     string
	DB           int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
}

func NewRedisClient() iRedis.Client {
	return &client{}
}

func (r *client) SetAddress(address string) iRedis.Client {
	r.Address = address
	return r
}

func (r *client) SetDatabase(database int) iRedis.Client {
	fmt.Println(r.Address)
	r.DB = database
	return r
}

func (r *client) SetPassword(password string) iRedis.Client {
	r.Password = password
	return r
}

func (r *client) SetTimeout(readTimeout, writeTimeout int) iRedis.Client {
	r.ReadTimeout = time.Duration(readTimeout) * time.Second
	r.WriteTimeout = time.Duration(writeTimeout) * time.Second
	return r
}

func (r *client) SetPoolSize(poolSize int) iRedis.Client {
	r.PoolSize = poolSize
	return r
}

func (r *client) Call() iRedis.SDK {
	return &sdk{
		redisClient: redis.NewClient(
			&redis.Options{
				Addr:         r.Address,
				Password:     r.Password,
				DB:           r.DB,
				ReadTimeout:  r.ReadTimeout,
				WriteTimeout: r.WriteTimeout,
				PoolSize:     r.PoolSize,
			},
		),
	}
}
