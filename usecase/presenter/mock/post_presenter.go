package mock

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"github.com/stretchr/testify/mock"
)

type PostPresenterMock struct {
	mock.Mock
}

func (m *PostPresenterMock) StorePostToModel(post *dto.StorePost) *model.Post {
	args := m.Called(post)
	return args.Get(0).(*model.Post)
}

func (m *PostPresenterMock) UpdatePostToModel(post *dto.Post) *model.Post {
	args := m.Called(post)
	return args.Get(0).(*model.Post)
}

func (m *PostPresenterMock) ResponsePost(post *model.Post) *dto.Post {
	args := m.Called(post)
	return args.Get(0).(*dto.Post)
}

func (m *PostPresenterMock) ResponsePosts(posts []*model.Post) []*dto.Post {
	args := m.Called(posts)
	return args.Get(0).([]*dto.Post)
}
