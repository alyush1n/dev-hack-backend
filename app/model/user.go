package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Clubs     []string
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
	Name       string             `json:"name" bson:"name" binding:"required"`
	Location   string             `json:"location" bson:"location" binding:"required"`
	Date       time.Time          `json:"date" bson:"date"`
	Attachment `json:"attachment" bson:"attachment"`
	IsUnique   bool
}

type Attachment struct {
	Id     primitive.ObjectID `bson:"_id"`
	URL    string             `json:"url" bson:"url" binding:"required"`
	SentBy string             `json:"sent_by" bson:"sent_by" binding:"required"`
}

type Clubs struct {
	IncomingEvents  []string
	Logo            Attachment
	BackgroundImage Attachment
}
