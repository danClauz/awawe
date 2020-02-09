package mock

import (
	"awawe/domain/dto"
	"context"
	"github.com/stretchr/testify/mock"
)

type UserInteractorMock struct {
	mock.Mock
}

func (m *UserInteractorMock) Store(ctx context.Context, user *dto.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *UserInteractorMock) FindAll(ctx context.Context) ([]*dto.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*dto.User), args.Error(1)
}

func (m *UserInteractorMock) GetByID(ctx context.Context, ID int) (*dto.User, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(*dto.User), args.Error(1)
}
