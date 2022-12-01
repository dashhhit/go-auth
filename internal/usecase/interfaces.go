package usecase

import (
	"auth-golang/internal/entity"
	"context"
)

type (
	UserRepo interface {
		AddUser(ctx context.Context, user entity.User) error
		IsExistsEmail(ctx context.Context, email string) bool
		GetUser(ctx context.Context, id string) (*entity.User, error)
	}
	UserHandlers interface {
		Create(ctx context.Context, user entity.User) error
		Get(ctx context.Context, id string) (*entity.User, error)
	}
)
