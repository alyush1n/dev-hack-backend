package db

import (
	"context"
	"dev-hack-backend/app/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateEvent(event model.Event) (isExist bool) {
	filter := bson.M{"_id": event.Id}
	_, err := eventsCollection.UpdateOne(context.Background(), filter, event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func UpdateUser(user model.User) (isExist bool) {
	filter := bson.M{"_id": user.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", user.Username},
			{"password", user.Password},
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"sex", user.Sex},
			{"points", user.Points},
		}},
		{"$push", bson.D{
			{"clubs", user.Clubs},
		}},
	}

	_, err := usersCollection.UpdateOne(context.Background(), update, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func AddEventToClub(clubName string, eventID primitive.ObjectID) {
	filter := bson.M{"name": clubName}
	update := bson.M{"$push": bson.M{"incoming_events": eventID}}

	_, err := clubsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}

func AddRegisteredEventToUser(username string, eventID primitive.ObjectID) {
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"registered_events": eventID}}

	_, err := usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}
