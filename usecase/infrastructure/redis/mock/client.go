package mock

import (
	"awawe/usecase/infrastructure/redis"
	"github.com/stretchr/testify/mock"
)

type RedisClientMock struct {
	mock.Mock
}

func (m *RedisClientMock) SetAddress(address string) redis.Client {
	args := m.Called(address)
	return args.Get(0).(redis.Client)
}

func (m *RedisClientMock) SetDatabase(database int) redis.Client {
	args := m.Called(database)
	return args.Get(0).(redis.Client)
}

func (m *RedisClientMock) SetPassword(password string) redis.Client {
	args := m.Called(password)
	return args.Get(0).(redis.Client)
}

func (m *RedisClientMock) SetTimeout(readTimeout, writeTimeout int) redis.Client {
	args := m.Called(readTimeout, writeTimeout)
	return args.Get(0).(redis.Client)
}

func (m *RedisClientMock) SetPoolSize(poolSize int) redis.Client {
	args := m.Called(poolSize)
	return args.Get(0).(redis.Client)
}

func (m *RedisClientMock) Call() redis.SDK {
	args := m.Called()
	return args.Get(0).(redis.SDK)
}
