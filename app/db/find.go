package db

import (
	"context"
	"dev-hack-backend/app/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetClubByName(name string) (club model.Club) {
	filter := bson.M{"name": name}

	err := clubsCollection.FindOne(context.Background(), filter).Decode(&club)
	if err != nil {
		fmt.Println(err)
		return model.Club{}
	}
	return club
}

func GetEventByID(Id primitive.ObjectID) (event model.Event, isExist bool) {
	filter := bson.M{"_id": Id}

	err := eventsCollection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		fmt.Println(err)
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
