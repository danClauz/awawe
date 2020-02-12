package mock

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type RedisSdkMock struct {
	mock.Mock
}

func (m *RedisSdkMock) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *RedisSdkMock) Set(key string, value interface{}, expiration time.Duration) error {
	args := m.Called(key, value, expiration)
	return args.Error(0)
}

func (m *RedisSdkMock) Get(key string) (interface{}, error) {
	args := m.Called(key)
	return args.Get(0).(interface{}), args.Error(1)
}
