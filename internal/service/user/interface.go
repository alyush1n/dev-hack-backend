package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	Service interface {
		GetUser(ctx context.Context) (*user.User, error)
		Authorize(ctx context.Context, username, password string) (*user.User, error)

		InsertUser(ctx context.Context, user *user.User) error
		UpdateUser(ctx context.Context, user *user.User) error
		DeleteUser(ctx context.Context) error

		CreateSession(ctx context.Context, id string) (string, string, error)
		RefreshToken(ctx context.Context, id, rToken string) (string, string, error)

		ParseToken(accessToken string) (string, error)
		ContextWithTimeout(c *gin.Context, ttl time.Duration, key string) (context.Context, context.CancelFunc)
		CompareID(ctx context.Context, compareUserID string) bool
	}

	JWTManager interface {
		NewJWT(userId string, ttl time.Duration) (string, error)
		NewRefreshToken() (string, error)
		ParseToken(accessToken string) (string, error)
	}

	Storage interface {
		GetUserById(ctx context.Context, id string) (*user.User, error)
		GetUserByUsername(ctx context.Context, username string) (*user.User, error)
		GetUserByRT(ctx context.Context, id string, rToken string) (*user.User, error)

		InsertUser(ctx context.Context, user *user.User) error

		UpdateUser(ctx context.Context, user *user.User) error
		UpdateSession(ctx context.Context, id string, session user.Session) error

		DeleteUserById(ctx context.Context, id string) error
	}
)
