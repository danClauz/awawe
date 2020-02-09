package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
)

type UserPresenter interface {
	RequestToModel(user *dto.User) *model.User
	ResponseUser(user *model.User) *dto.User
	ResponseUsers(users []*model.User) []*dto.User
}
