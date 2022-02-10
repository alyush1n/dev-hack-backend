package user

import (
	"context"
	"dev-hack-backend/internal/adapters/api/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type service struct {
	storage         Storage
	jwt             JWTManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewService(storage Storage, jwt JWTManager, accessTokenTTL, refreshTokenTTL time.Duration) user.Service {
	return &service{
		storage:         storage,
		jwt:             jwt,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *service) GetUserById(ctx context.Context, id string) (*User, error) {
	return s.storage.GetUserById(ctx, id)
}

func (s *service) InsertUser(ctx context.Context, dto *user.CreateUserDTO) (*User, error) {
	return nil, nil
}

func (s *service) UpdateUser(ctx context.Context, dto *user.UpdateUserDTO) (*User, error) {
	return nil, nil
}

func (s *service) DeleteUserById(ctx context.Context, id string) error {
	return nil
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
	session := Session{RefreshToken: rToken, ExpiresAt: time.Now().Add(s.refreshTokenTTL)}
	err = s.storage.UpdateSession(ctx, id, session)
	if err != nil {
		return "", "", err
	}
	return aToken, rToken, nil
}

func (s *service) RefreshToken(ctx context.Context, id primitive.ObjectID, rToken string) (string, string, error) {
	user, err := s.storage.GetUserByRT(ctx, id.Hex(), rToken)
	if err != nil {
		return "", "", err
	}

	return s.CreateSession(ctx, user.Id.Hex())
}
