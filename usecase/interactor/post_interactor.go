package interactor

import (
	"awawe/domain/dto"
	"awawe/usecase/presenter"
	"awawe/usecase/repository"
	"context"
)

type postInteractor struct {
	postRepository repository.PostRepository
	postPresenter  presenter.PostPresenter
}

type PostInteractor interface {
	Store(ctx context.Context, post *dto.Post) error
	Update(ctx context.Context, post *dto.Post) error
	FindAll(ctx context.Context) ([]*dto.Post, error)
}

func NewPostInteractor(r repository.PostRepository, p presenter.PostPresenter) PostInteractor {
	return &postInteractor{
		postRepository: r,
		postPresenter:  p,
	}
}

func (in *postInteractor) Store(ctx context.Context, post *dto.Post) error {
	if err := in.postRepository.Store(ctx, in.postPresenter.RequestToModel(post)); err != nil {
		return err
	}

	return nil
}

func (in *postInteractor) Update(ctx context.Context, post *dto.Post) error {
	if err := in.postRepository.Update(ctx, in.postPresenter.RequestToModel(post)); err != nil {
		return err
	}

	return nil
}

func (in *postInteractor) FindAll(ctx context.Context) ([]*dto.Post, error) {
	posts, err := in.postRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return in.postPresenter.ResponsePosts(posts), nil
}
