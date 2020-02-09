package mock

import (
	"awawe/domain/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Store(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindAll(ctx context.Context) ([]*model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByID(ctx context.Context, ID int) (*model.User, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(*model.User), args.Error(1)
}
