package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Clubs     []string           `json:"clubs" bson:"clubs"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Sex       string             `json:"sex" bson:"sex"`
	Points    string             `json:"points" bson:"points"`
}

type Event struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Type        string             `json:"type" bson:"type"`
	Clubs       []string           `json:"clubs" bson:"clubs"`
	Count       int               `json:"count" bson:"count"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Location    string             `json:"location" bson:"location"`
	Date        string             `json:"date" bson:"date"`
	Attachment  `json:"attachment" bson:"attachment"`
}

type Attachment struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	URL    string             `json:"url" bson:"url" binding:"required"`
	SentBy string             `json:"sent_by" bson:"sent_by" binding:"required"`
}

type Club struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Count           int               `json:"count" bson:"count"`
	IncomingEvents  []string           `json:"incoming_events" bson:"incoming_events"`
	Logo            Attachment         `json:"logo" bson:"logo"`
	BackgroundImage Attachment         `json:"background_image" bson:"background_image"`
}
