package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
)

type Service interface {
	GetUserById(ctx context.Context, id string) (*user.User, error)
	GetUserByUsername(ctx context.Context, dto *SignInUserDTO) (*user.User, error)
	InsertUser(ctx context.Context, dto *CreateUserDTO) (*user.User, error)
	UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*user.User, error)
	DeleteUserById(ctx context.Context, id string) error

	CreateSession(ctx context.Context, id string) (string, string, error)
	RefreshToken(ctx context.Context, id, rToken string) (string, string, error)

	HashPassword(ctx context.Context, password string) string
}
