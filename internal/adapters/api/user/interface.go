package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
)

type Service interface {
	GetUser(ctx context.Context) (*user.User, error)
	Authorize(ctx context.Context, dto *SignInUserDTO) (*user.User, error)

	InsertUser(ctx context.Context, dto *CreateUserDTO) (*user.User, error)
	UpdateUser(ctx context.Context, dto *UpdateUserDTO) (*user.User, error)
	DeleteUser(ctx context.Context) error

	CreateSession(ctx context.Context, id string) (string, string, error)
	RefreshToken(ctx context.Context, id, rToken string) (string, string, error)

	ParseToken(accessToken string) (string, error)
}
