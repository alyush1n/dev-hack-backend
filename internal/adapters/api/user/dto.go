package user

import "go.mongodb.org/mongo-driver/bson/primitive"

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
