package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"awawe/usecase/presenter"
)

type postPresenter struct{}

func NewPostPresenter() presenter.PostPresenter {
	return &postPresenter{}
}

func (pre *postPresenter) RequestToModel(post *dto.Post) *model.Post {
	return &model.Post{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}

func (pre *postPresenter) ResponsePost(post *model.Post) *dto.Post {
	return &dto.Post{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}

func (pre *postPresenter) ResponsePosts(posts []*model.Post) []*dto.Post {
	response := make([]*dto.Post, 0)

	for _, post := range posts {
		tmp := pre.ResponsePost(post)
		response = append(response, tmp)
	}

	return response
}
