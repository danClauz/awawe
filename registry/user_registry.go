package registry

import (
	config "awawe/configuration"
	ip "awawe/delivery/presenter"
	ir "awawe/delivery/repository"
	"awawe/infrastructure/redis"
	"awawe/usecase/interactor"
	up "awawe/usecase/presenter"
	ur "awawe/usecase/repository"
	"database/sql"
)

func NewUserInteractor(db *sql.DB) interactor.UserInteractor {
	return interactor.NewUserInteractor(newUserRepository(db), newUserPresenter())
}

func newUserRepository(db *sql.DB) ur.UserRepository {
	redisConfig := config.GetRedisConfig()
	return ir.NewUserRepository(db, redis.NewRedisClient().
		SetAddress(redisConfig.Address).
		SetDatabase(redisConfig.DB).
		SetPassword(redisConfig.Password).
		SetTimeout(redisConfig.ReadTimeout, redisConfig.WriteTimeout).
		SetPoolSize(redisConfig.PoolSize).
		Call())
}

func newUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
