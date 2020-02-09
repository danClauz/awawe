package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
)

type PostPresenter interface {
	RequestToModel(post *dto.Post) *model.Post
	ResponsePost(post *model.Post) *dto.Post
	ResponsePosts(posts []*model.Post) []*dto.Post
}
