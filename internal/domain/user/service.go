package user

import (
	"context"
	"crypto/sha1"
	"dev-hack-backend/internal/adapters/api/user"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	ComparePassError = "password incorrect"
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

func (s *service) GetUserByUsername(ctx context.Context, dto *user.SignInUserDTO) (*User, error) {
	user, err := s.storage.GetUserByUsername(ctx, dto.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != s.HashPassword(ctx, dto.Password) {
		return nil, fmt.Errorf(ComparePassError)
	}
	return user, nil
}

func (s *service) InsertUser(ctx context.Context, dto *user.CreateUserDTO) (*User, error) {
	user := &User{
		Id:            primitive.NewObjectID(),
		Username:      dto.Username,
		Password:      dto.Password,
		PhotoURL:      "",
		Clubs:         nil,
		VisitedEvents: nil,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Sex:           dto.Sex,
		Session: Session{
			RefreshToken: "",
			ExpiresAt:    time.Time{},
		},
		Points: 0,
	}
	return s.storage.InsertUser(ctx, user)
}

func (s *service) UpdateUser(ctx context.Context, dto *user.UpdateUserDTO) (*User, error) {
	user := &User{
		Id:            primitive.NewObjectID(),
		Username:      dto.Username,
		Password:      dto.Password,
		PhotoURL:      dto.PhotoURL,
		Clubs:         dto.Clubs,
		VisitedEvents: dto.VisitedEvents,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Sex:           dto.Sex,
		Session:       dto.Session,
		Points:        dto.Points,
	}
	return s.storage.UpdateUser(ctx, user)
}

func (s *service) DeleteUserById(ctx context.Context, id string) error {
	return s.storage.DeleteUser(ctx, id)
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

func (s *service) RefreshToken(ctx context.Context, id, rToken string) (string, string, error) {
	_, err := s.storage.GetUserByRT(ctx, id, rToken)
	if err != nil {
		return "", "", err
	}

	return s.CreateSession(ctx, id)
}

func (s *service) HashPassword(ctx context.Context, password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", password)
}
