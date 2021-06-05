package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteEvent(event model.Event) (isExist bool) {
	filter := bson.M{"_id": event.Id}

	_, err := eventsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		return
	}
	return true
}
