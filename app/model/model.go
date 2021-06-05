package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID   `bson:"_id"`
	Username  string               `json:"username" bson:"username" binding:"required"`
	Password  string               `json:"password" bson:"password" binding:"required"`
	Clubs     []primitive.ObjectID `json:"clubs" bson:"clubs"` //tags
	FirstName string               `json:"first_name" bson:"first_name"`
	LastName  string               `json:"last_name" bson:"last_name"`
	Sex       string               `json:"sex" bson:"sex"`
	Points    string               `json:"points" bson:"points"`
	City      string               `json:"city" bson:"city"`
	Stats
}

type Stats struct {
}

type Event struct {
	Id              primitive.ObjectID `bson:"_id"`
	Name            string             `json:"name" bson:"name" binding:"required"`
	City            string             `json:"city" bson:"city" binding:"required"`
	Date            string             `json:"date" bson:"date"`
	Logo            Attachment         `json:"logo" bson:"logo"`
	BackgroundImage Attachment         `json:"background_image" bson:"background_image"` //обратимся к SendBy от фотокарточки и получим отправителя (создателя ивента)
}

type Attachment struct {
	URL    string `json:"url" bson:"url" binding:"required"`
	SentBy string `json:"sent_by" bson:"sent_by" binding:"required"`
}

type Club struct { //tags
	Id     primitive.ObjectID   `bson:"_id"`
	Events []primitive.ObjectID `json:"events" bson:"events"`
	Logo   Attachment           `json:"logo" bson:"logo"`
}
