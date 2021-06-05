package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func eventCollection() (collection *mongo.Collection) {
	collection = client.Database("dev-hack").Collection("events")
	return collection
}

func InsertEvent(Event model.Event) (isExist bool) {
	Event.Id = primitive.NewObjectID()
	_, err := eventCollection().InsertOne(context.Background(), Event)
	if err != nil {
		return false
	}
	return true
}

func FindEventById(Id primitive.ObjectID) (event model.Event, isExist bool) {
	filter := bson.M{"_id": Id}

	err := eventCollection().FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Event{}, false
		}
		return
	}
	return event, true
}

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

	_, err := eventCollection().UpdateOne(context.Background(), update, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		return
	}

	return true
}

func DeleteEvent(event model.Event) (isExist bool) {
	filter := bson.M{"_id": event.Id}

	_, err := eventCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		return
	}
	return true
}
