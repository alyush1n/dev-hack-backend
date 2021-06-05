package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Clubs     []primitive.ObjectID
	FirstName string
	LastName  string
	Sex       string
	Points    string
	Stats
}

type Stats struct {
}

type Event struct {
	Id         primitive.ObjectID `bson:"_id"`
	Clubs      []string           `json:"clubs" bson:"clubs"`
	Type       string             `json:"type"`
	Name       string             `json:"name" bson:"name"`
	Location   string             `json:"location" bson:"location"`
	Date       string             `json:"date" bson:"date"`
	Attachment `json:"attachment" bson:"attachment"`
}

type Attachment struct {
	Id     primitive.ObjectID `bson:"_id"`
	URL    string             `json:"url" bson:"url" binding:"required"`
	SentBy string             `json:"sent_by" bson:"sent_by" binding:"required"`
}

type Club struct {
	Id              primitive.ObjectID `bson:"_id"`
	Name            string
	IncomingEvents  []primitive.ObjectID
	Logo            Attachment
	BackgroundImage Attachment
}
