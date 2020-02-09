package registry

import (
	"awawe/infrastucture/redis"
	ip "awawe/interface/presenter"
	ir "awawe/interface/repository"
	"awawe/usecase/interactor"
	up "awawe/usecase/presenter"
	ur "awawe/usecase/repository"
	"database/sql"
)

func NewUserInteractor(db *sql.DB) interactor.UserInteractor {
	return interactor.NewUserInteractor(newUserRepository(db), newUserPresenter())
}

func newUserRepository(db *sql.DB) ur.UserRepository {
	return ir.NewUserRepository(db, redis.NewRedisClient())
}

func newUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
