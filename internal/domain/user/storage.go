package user

import (
	"context"
)

type Storage interface {
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByRT(ctx context.Context, id string, rToken string) (*User, error)

	InsertUser(ctx context.Context, user *User) (*User, error)

	UpdateUser(ctx context.Context, user *User) (*User, error)
	UpdateSession(ctx context.Context, id string, session Session) error

	DeleteUserById(ctx context.Context, id string) error
}
