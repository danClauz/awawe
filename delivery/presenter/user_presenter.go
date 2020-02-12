package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"awawe/usecase/presenter"
)

type userPresenter struct {
	post presenter.PostPresenter
}

func NewUserPresenter() presenter.UserPresenter {
	post := new(postPresenter)
	return &userPresenter{
		post: post,
	}
}

func (pre *userPresenter) RequestToModel(user *dto.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func (pre *userPresenter) ResponseUser(user *model.User) *dto.User {
	response := &dto.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	if user.Posts != nil {
		response.Posts = pre.post.ResponsePosts(user.Posts)
	}

	return response
}

func (pre *userPresenter) ResponseUsers(users []*model.User) []*dto.User {
	response := make([]*dto.User, 0)

	for _, user := range users {
		tmp := pre.ResponseUser(user)
		response = append(response, tmp)
	}

	return response
}
