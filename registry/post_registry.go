package registry

import (
	ip "awawe/delivery/presenter"
	ir "awawe/delivery/repository"
	"awawe/usecase/interactor"
	up "awawe/usecase/presenter"
	ur "awawe/usecase/repository"
	"database/sql"
)

func NewPostInteractor(db *sql.DB) interactor.PostInteractor {
	return interactor.NewPostInteractor(newPostRepository(db), newUserRepository(db), newPostPresenter())
}

func newPostRepository(db *sql.DB) ur.PostRepository {
	return ir.NewPostRepository(db)
}

func newPostPresenter() up.PostPresenter {
	return ip.NewPostPresenter()
}
