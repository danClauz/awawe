package registry

import (
	ip "awawe/interface/presenter"
	ir "awawe/interface/repository"
	"awawe/usecase/interactor"
	up "awawe/usecase/presenter"
	ur "awawe/usecase/repository"
	"database/sql"
)

func NewPostInteractor(db *sql.DB) interactor.PostInteractor {
	return interactor.NewPostInteractor(newPostRepository(db), newPostPresenter())
}

func newPostRepository(db *sql.DB) ur.PostRepository {
	return ir.NewPostRepository(db)
}

func newPostPresenter() up.PostPresenter {
	return ip.NewPostPresenter()
}
