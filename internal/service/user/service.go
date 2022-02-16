package user

import (
	"context"
	"crypto/sha1"
	user "dev-hack-backend/internal/domain/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	ComparePassError = "password incorrect"
	ContextKey       = "user_id"
)

type service struct {
	storage         Storage
	jwt             JWTManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewService(storage Storage, jwt JWTManager, accessTokenTTL, refreshTokenTTL time.Duration) Service {
	return &service{
		storage:         storage,
		jwt:             jwt,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *service) GetUser(ctx context.Context) (*user.User, error) {
	userId := ctx.Value(ContextKey)

	return s.storage.GetUserById(ctx, fmt.Sprintf("%s", userId))
}

func (s *service) Authorize(ctx context.Context, username, password string) (*user.User, error) {
	currentUser, err := s.storage.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if currentUser.Password != s.HashPassword(password) {
		return nil, fmt.Errorf(ComparePassError)
	}

	return currentUser, nil
}

func (s *service) InsertUser(ctx context.Context, currentUser *user.User) error {
	currentUser.Password = s.HashPassword(currentUser.Password)
	return s.storage.InsertUser(ctx, currentUser)
}

func (s *service) UpdateUser(ctx context.Context, currentUser *user.User) error {
	return s.storage.UpdateUser(ctx, currentUser)
}

func (s *service) DeleteUser(ctx context.Context) error {
	userId := ctx.Value(ContextKey)

	return s.storage.DeleteUserById(ctx, fmt.Sprintf("%s", userId))
}

func (s *service) CreateSession(ctx context.Context, id string) (string, string, error) {
	aToken, err := s.jwt.NewJWT(id, s.accessTokenTTL)
	if err != nil {
		return "", "", err
	}

	rToken, err := s.jwt.NewRefreshToken()
	if err != nil {
		return "", "", err
	}

	session := user.Session{RefreshToken: rToken, ExpiresAt: time.Now().Add(s.refreshTokenTTL)}
	err = s.storage.UpdateSession(ctx, id, session)
	if err != nil {
		return "", "", err
	}

	return aToken, rToken, nil
}

func (s *service) RefreshToken(ctx context.Context, id, rToken string) (string, string, error) {
	_, err := s.storage.GetUserByRT(ctx, id, rToken)
	if err != nil {
		return "", "", err
	}

	return s.CreateSession(ctx, id)
}

func (s *service) HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", password)
}

func (s *service) ParseToken(accessToken string) (string, error) {
	return s.jwt.ParseToken(accessToken)
}

func (s *service) CompareID(ctx context.Context, compareUserID string) bool {
	userID := fmt.Sprintf("%s", ctx.Value(ContextKey))
	if userID != compareUserID {
		return false
	}

	return true
}

func (s *service) ContextWithTimeout(c *gin.Context, ttl time.Duration, key string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), ttl)
	ctxValue, ok := c.Get(ContextKey)
	if !ok {
		return ctx, cancel
	}
	ctx = context.WithValue(ctx, ContextKey, ctxValue)
	return ctx, cancel
}
