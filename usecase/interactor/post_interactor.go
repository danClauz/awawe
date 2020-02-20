package interactor

import (
	"awawe/domain/dto"
	"awawe/usecase/presenter"
	"awawe/usecase/repository"
	"context"
)

type postInteractor struct {
	postRepository repository.PostRepository
	userRepository repository.UserRepository
	postPresenter  presenter.PostPresenter
}

type PostInteractor interface {
	Store(ctx context.Context, post *dto.StorePost) error
	Update(ctx context.Context, post *dto.Post) error
	FindAll(ctx context.Context) ([]*dto.Post, error)
}

func NewPostInteractor(r repository.PostRepository, userRepository repository.UserRepository, p presenter.PostPresenter) PostInteractor {
	return &postInteractor{
		postRepository: r,
		userRepository: userRepository,
		postPresenter:  p,
	}
}

func (in *postInteractor) Store(ctx context.Context, post *dto.StorePost) error {
	if err := in.postRepository.Store(ctx, in.postPresenter.StorePostToModel(post)); err != nil {
		return err
	}

	return nil
}

func (in *postInteractor) Update(ctx context.Context, post *dto.Post) error {
	if err := in.postRepository.Update(ctx, in.postPresenter.UpdatePostToModel(post)); err != nil {
		return err
	}

	return nil
}

func (in *postInteractor) FindAll(ctx context.Context) ([]*dto.Post, error) {
	posts, err := in.postRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		if post.User, err = in.userRepository.GetByID(ctx, int(post.UserID)); err != nil {
			return nil, err
		}
	}

	return in.postPresenter.ResponsePosts(posts), nil
}
