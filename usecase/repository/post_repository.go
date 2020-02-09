package repository

import (
	"awawe/domain/model"
	"context"
)

type PostRepository interface {
	Store(ctx context.Context, post *model.Post) error
	Update(ctx context.Context, post *model.Post) error
	FindAll(ctx context.Context) ([]*model.Post, error)
}
