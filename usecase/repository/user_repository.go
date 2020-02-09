package repository

import (
	"awawe/domain/model"
	"context"
)

type UserRepository interface {
	StoreToRedis(ctx context.Context, user *model.User) error
	FindAll(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, ID int) (*model.User, error)
}
