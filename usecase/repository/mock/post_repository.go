package mock

import (
	"awawe/domain/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type PostRepositoryMock struct {
	mock.Mock
}

func (m *PostRepositoryMock) Store(ctx context.Context, post *model.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *PostRepositoryMock) Update(ctx context.Context, post *model.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *PostRepositoryMock) FindAll(ctx context.Context) ([]*model.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Post), args.Error(1)
}
