package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func usersCollection() (collection *mongo.Collection) {
	collection = client.Database("dev-hack").Collection("users")
	return collection
}

func InsertUser(User model.User) (err error) {
	_, err = usersCollection().InsertOne(context.Background(), User)
	if err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(Username string) (User model.User, isExist bool) {
	filter := bson.M{"username": Username}

	err := usersCollection().FindOne(context.Background(), filter).Decode(&User)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return User, true
}

func UpdateUser(User model.User) (isExist bool) {
	filter := bson.M{"_id": User.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", User.Username},
			{"password", User.Password},
		}},
	}

	_, err := usersCollection().UpdateOne(context.Background(), update, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}

func DeleteUser(User model.User) (isExist bool) {
	filter := bson.M{"_id": User.Id}

	_, err := usersCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		return
	}

	return true
}
