package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetClubByID(Id primitive.ObjectID) (club model.Club) {
	filter := bson.M{"_id": Id}

	err := clubsCollection.FindOne(context.Background(), filter).Decode(&club)
	if err != nil {
		return model.Club{}
	}
	return club
}

func GetEventByID(Id primitive.ObjectID) (event model.Event, isExist bool) {
	filter := bson.M{"_id": Id}

	err := eventCollection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Event{}, false
		}
		return
	}
	return event, true
}

func FindUserByUsername(Username string) (User model.User, isExist bool) {
	filter := bson.M{"username": Username}

	err := usersCollection.FindOne(context.Background(), filter).Decode(&User)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return User, true
}
