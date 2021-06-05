package db

import (
	"context"
	"dev-hack-backend/app/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client           *mongo.Client
	usersCollection  *mongo.Collection
	clubsCollection  *mongo.Collection
	eventsCollection *mongo.Collection
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://fizik:"+config.MongoPass+"@cluster0.oeuni.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("dev-hack")
	usersCollection = database.Collection("users")
	clubsCollection = database.Collection("clubs")
	eventsCollection = database.Collection("events")

	fmt.Println("Connected to MongoDB!")
}
