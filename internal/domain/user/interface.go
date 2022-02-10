package user

import (
	"context"
	"time"
)

type Storage interface {
	GetUserById(ctx context.Context, id string) (*User, error)
	GetUserByRT(ctx context.Context, id string, rToken string) (*User, error)
	InsertUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	UpdateSession(ctx context.Context, id string, session Session) error
	DeleteUser(ctx context.Context, id string) error
}

type JWTManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
	ParseToken(accessToken string) (string, error)
}
