package mongodb

import (
	"context"
	"dev-hack-backend/pkg/apperror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewClient(ctx context.Context, mongoURI, database string) (*mongo.Database, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	log.Println("mongo connect")
	option := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(c, option)
	if err != nil {
		return nil, err
	}

	log.Println("mongo ping")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, apperror.ClientError
	}
	log.Println("connected to mongo")

	return client.Database(database), nil
}
