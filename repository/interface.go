package repository

import (
	"context"

	"github.com/techierishi/mongocrud/model"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, in model.User) (model.User, error)
	UpdateUser(ctx context.Context, in model.User) (model.User, error)
	DeleteUser(ctx context.Context, email string) error
}
