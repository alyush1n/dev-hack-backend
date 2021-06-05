package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func clubsCollection() (collection *mongo.Collection) {
	collection = client.Database("dev-hack").Collection("clubs")
	return collection
}

func GetClubByID(Id primitive.ObjectID) (club model.Club) {
	filter := bson.M{"_id": Id}

	err := clubsCollection().FindOne(context.Background(), filter).Decode(&club)
	if err != nil {
		return model.Club{}
	}
	return club
}
