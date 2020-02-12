package interactor

import (
	"awawe/domain/dto"
	"awawe/usecase/presenter"
	"awawe/usecase/repository"
	"context"
)

type userInteractor struct {
	userRepository repository.UserRepository
	userPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Store(ctx context.Context, user *dto.User) error
	StoreToRedis(ctx context.Context, user *dto.User) error
	FindAll(ctx context.Context) ([]*dto.User, error)
	GetByID(ctx context.Context, ID int) (*dto.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{
		userRepository: r,
		userPresenter:  p,
	}
}

func (in *userInteractor) Store(ctx context.Context, user *dto.User) error {
	if err := in.userRepository.Store(ctx, in.userPresenter.RequestToModel(user)); err != nil {
		return err
	}

	return nil
}

func (in *userInteractor) StoreToRedis(ctx context.Context, user *dto.User) error {
	if err := in.userRepository.StoreToRedis(ctx, in.userPresenter.RequestToModel(user)); err != nil {
		return err
	}

	return nil
}

func (in *userInteractor) FindAll(ctx context.Context) ([]*dto.User, error) {
	user, err := in.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return in.userPresenter.ResponseUsers(user), nil
}

func (in *userInteractor) GetByID(ctx context.Context, ID int) (*dto.User, error) {
	user, err := in.userRepository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return in.userPresenter.ResponseUser(user), nil
}
