package presenter

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"awawe/usecase/presenter"
)

type postPresenter struct {
	user presenter.UserPresenter
}

func NewPostPresenter() presenter.PostPresenter {
	user := new(userPresenter)
	return &postPresenter{
		user: user,
	}
}

func (pre *postPresenter) StorePostToModel(post *dto.StorePost) *model.Post {
	return &model.Post{
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}

func (pre *postPresenter) UpdatePostToModel(post *dto.Post) *model.Post {
	return &model.Post{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
	}
}

func (pre *postPresenter) ResponsePost(post *model.Post) *dto.Post {
	response := &dto.Post{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	if post.User != nil {
		response.User = pre.user.ResponseUser(post.User)
	}

	return response
}

func (pre *postPresenter) ResponsePosts(posts []*model.Post) []*dto.Post {
	response := make([]*dto.Post, 0)

	for _, post := range posts {
		tmp := pre.ResponsePost(post)
		response = append(response, tmp)
	}

	return response
}
