package redis

type Client interface {
	SetAddress(address string) Client
	SetDatabase(database int) Client
	SetPassword(password string) Client
	SetTimeout(readTimeout, writeTimeout int) Client
	SetPoolSize(poolSize int) Client
	Call() SDK
}
