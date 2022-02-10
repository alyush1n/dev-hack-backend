package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
)

type Service interface {
	GetUserById(ctx context.Context, id string) (*user.User, error)
	InsertUser(ctx context.Context, dto *CreateUserDTO) (*user.User, error)
	UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*user.User, error)
	DeleteUserById(ctx context.Context, id string) error
}
