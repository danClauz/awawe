package mock

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"github.com/stretchr/testify/mock"
)

type UserPresenterMock struct {
	mock.Mock
}

func (m *UserPresenterMock) RequestToModel(user *dto.User) *model.User {
	args := m.Called(user)
	return args.Get(0).(*model.User)
}

func (m *UserPresenterMock) ResponseUser(user *model.User) *dto.User {
	args := m.Called(user)
	return args.Get(0).(*dto.User)
}

func (m *UserPresenterMock) ResponseUsers(users []*model.User) []*dto.User {
	args := m.Called(users)
	return args.Get(0).([]*dto.User)
}
