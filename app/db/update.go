package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateEvent(event model.Event) (isExist bool) {
	filter := bson.M{"_id": event.Id}

	update := bson.D{
		{"$set", bson.D{
			{"name", event.Name},
			{"location", event.Location},
			{"date", event.Date},
			{"attachment", bson.D{
				{"url", event.URL},
				{"sent_by", event.SentBy},
			}},
		}},
	}

	_, err := eventCollection.UpdateOne(context.Background(), update, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		return
	}

	return true
}

func UpdateUser(User model.User) (isExist bool) {
	filter := bson.M{"_id": User.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", User.Username},
			{"password", User.Password},
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
