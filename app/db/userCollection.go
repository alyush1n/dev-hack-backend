package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func usersCollection() (collection *mongo.Collection) {
	collection = client.Database("dev-hack").Collection("users")
	return collection
}

func InsertUser(User model.User) (err error) {
	User.Id = primitive.NewObjectID()
	_, err = usersCollection().InsertOne(context.Background(), User)
	if err != nil {
		return err
	}
	return nil
}

func FindUserByID(ID string) (User model.User, isExist bool) {
	filter := bson.M{"id": ID}

	err := usersCollection().FindOne(context.Background(), filter).Decode(&User)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return User, true
}

func UpdateUser(User model.User) (err error) {
	filter := bson.M{"_id": User.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", User.Username},
			{"password", User.Password},
		}},
	}

	_, err = usersCollection().UpdateOne(context.Background(), update, filter)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(User model.User) (err error) {
	filter := bson.M{"_id": User.Id}

	_, err = usersCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
