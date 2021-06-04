package db

import (
	"context"
	"dev-hack-backend/app/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

func Connect() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://fizik:dKUAhJHSxBc3JS38zSxN@cluster0.oeuni.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func IsnsertUser(User model.User) (err error) {
	User.Id = primitive.NewObjectID()
	_, err = usersCollection().InsertOne(context.Background(), User)
	if err != nil {
		return err
	}
	return nil
}

func FindUserById(Id primitive.ObjectID) (User model.User, err error) {
	filter := bson.M{"_id": Id}

	err = usersCollection().FindOne(context.Background(), filter).Decode(&User)
	if err != nil {
		return model.User{}, err
	}
	return User, nil
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

func usersCollection() (collection *mongo.Collection) {
	collection = client.Database("dev-hack").Collection("users")
	return collection
}
