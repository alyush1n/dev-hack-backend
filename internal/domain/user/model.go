package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Username      string             `json:"username" bson:"username" binding:"required"`
	Password      string             `json:"password" bson:"password" binding:"required"`
	PhotoURL      string             `json:"photo_url" bson:"photo_url"`
	Clubs         []string           `json:"clubs" bson:"clubs"`
	VisitedEvents []string           `json:"visited_events" bson:"visited_events"`
	FirstName     string             `json:"first_name" bson:"first_name"`
	LastName      string             `json:"last_name" bson:"last_name"`
	Sex           string             `json:"sex" bson:"sex"`
	Session       Session            `json:"session" bson:"session"`
	Points        int                `json:"points" bson:"points"`
}

type Session struct {
	RefreshToken string
	ExpiresAt    time.Time
}
