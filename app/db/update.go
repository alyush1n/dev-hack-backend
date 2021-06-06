package db

import (
	"context"
	"dev-hack-backend/app/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//func UpdateEvent(event model.Event) (isExist bool) {
//	filter := bson.M{"_id": event.Id}
//
//	update := bson.D{
//		{"$set", bson.D{
//			{"name", event.Name},
//			{"location", event.City},
//			{"date", event.Date},
//			{"attachment", bson.D{
//				{"url", },
//				{"sent_by", event.SentBy},
//			}},
//		}},
//	}
//
//	_, err := eventsCollection.UpdateOne(context.Background(), update, filter)
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return false
//		}
//		return
//	}
//
//	return true
//}

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

func UpdateUser(User model.User) (isExist bool) {
	filter := bson.M{"_id": User.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", User.Username},
			{"password", User.Password},
			{"first_name", User.FirstName},
			{"last_name", User.LastName},
			{"sex", User.Sex},
			{"points", User.Points},
		}},
		{"$push", bson.D{
			{"clubs", User.Clubs},
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
