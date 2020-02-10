package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"awawe/usecase/presenter"
)

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (pre *userPresenter) RequestToModel(user *dto.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func (pre *userPresenter) ResponseUser(user *model.User) *dto.User {
	return &dto.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func (pre *userPresenter) ResponseUsers(users []*model.User) []*dto.User {
	response := make([]*dto.User, 0)

	for _, user := range users {
		tmp := pre.ResponseUser(user)
		response = append(response, tmp)
	}

	return response
}
