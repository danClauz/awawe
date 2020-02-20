package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
)

type PostPresenter interface {
	StorePostToModel(post *dto.StorePost) *model.Post
	UpdatePostToModel(post *dto.Post) *model.Post
	ResponsePost(post *model.Post) *dto.Post
	ResponsePosts(posts []*model.Post) []*dto.Post
}
