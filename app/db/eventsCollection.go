package db

import (
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func eventCollection() (collection *mongo.Collection) {
	collection = client.Database("dev-hack").Collection("events")
	return collection
}

func InsertEvent(Event model.Events) (err error) {
	//filter := bson.M{}
	return nil
}
