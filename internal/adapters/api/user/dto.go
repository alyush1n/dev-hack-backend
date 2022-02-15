package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateUserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       string `json:"sex"`
}

type UpdateUserDTO struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Username      string             `json:"username" `
	Password      string             `json:"password" `
	PhotoURL      string             `json:"photo_url" `
	Clubs         []string           `json:"clubs" `
	VisitedEvents []string           `json:"visited_events" `
	FirstName     string             `json:"first_name" `
	LastName      string             `json:"last_name" `
	Sex           string             `json:"sex" `
	Points        int                `json:"points" `
}

type SignInUserDTO struct {
	Username string
	Password string
}

func (dto *CreateUserDTO) toUser() *user.User {
	return &user.User{
		Id:            primitive.NewObjectID(),
		Username:      dto.Username,
		Password:      dto.Password,
		PhotoURL:      "",
		Clubs:         nil,
		VisitedEvents: nil,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Sex:           dto.Sex,
		Session: user.Session{
			RefreshToken: "",
			ExpiresAt:    time.Time{},
		},
		Points: 0,
	}
}

func (dto *UpdateUserDTO) toUser(ctx context.Context) (*user.User, error) {
	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%s", ctx.Value("user_id")))
	if err != nil {
		return nil, err
	}

	return &user.User{
		Id:            id,
		Username:      dto.Username,
		Password:      dto.Password,
		PhotoURL:      dto.PhotoURL,
		Clubs:         dto.Clubs,
		VisitedEvents: dto.VisitedEvents,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Sex:           dto.Sex,
		Points:        dto.Points,
	}, nil
}
