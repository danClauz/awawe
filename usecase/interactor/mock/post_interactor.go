package mock

import (
	"awawe/domain/dto"
	"context"
	"github.com/stretchr/testify/mock"
)

type PostInteractorMock struct {
	mock.Mock
}

func (m *PostInteractorMock) Store(ctx context.Context, post *dto.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *PostInteractorMock) Update(ctx context.Context, post *dto.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *PostInteractorMock) FindAll(ctx context.Context) ([]*dto.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*dto.Post), args.Error(1)
}
