package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" bson:"username" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}

type Events struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Location   string             `json:"address" bson:"address" binding:"required"`
	Date       time.Time          `json:"date" bson:"date"`
	Attachment string             `json:"attachment" bson:"attachment"`
}
