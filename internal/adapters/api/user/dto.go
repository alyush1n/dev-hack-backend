package user

import (
	"dev-hack-backend/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	contextKey = "user_id"
)

type CreateUserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       string `json:"sex"`
}

type ResponseUserDTO struct {
	Id        string       `json:"id"`
	Username  string       `json:"username"`
	PhotoURL  string       `json:"photoURL"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Sex       string       `json:"sex"`
	Session   user.Session `json:"session"`
}

type SignInUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDTO struct {
	RToken string `json:"refresh_token"`
}

func (dto *CreateUserDTO) toUser() *user.User {
	return &user.User{
		Id:        primitive.NewObjectID(),
		Username:  dto.Username,
		Password:  dto.Password,
		PhotoURL:  "",
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Sex:       dto.Sex,
		Session: user.Session{
			RefreshToken: "",
			ExpiresAt:    time.Time{},
		},
	}
}

func toDTO(user *user.User) *ResponseUserDTO {
	return &ResponseUserDTO{
		Id:        user.Id.Hex(),
		Username:  user.Username,
		PhotoURL:  user.PhotoURL,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Sex:       user.Sex,
		Session:   user.Session,
	}
}
