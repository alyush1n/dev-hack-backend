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

var client *mongo.Client

func Connect() {

	//var err error
	//
	//ctx := context.Background()
	//clientOptions := options.Client().ApplyURI("mongodb+srv://fizik:"+config.MongoPass+"@cluster0.oeuni.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	//client, err := mongo.NewClient(clientOptions)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//
	//err = client.Connect(ctx)
	//if err != nil {
	//	fmt.Println("Client error: "+err.Error())
	//}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://fizik:"+config.MongoPass+"@cluster0.oeuni.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
