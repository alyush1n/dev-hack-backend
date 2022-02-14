package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	clientError = "failed to create client to mongodb with error %w"
)

func NewClient(ctx context.Context, mongoURI, database string) (*mongo.Database, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf(clientError, err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf(clientError, err)
	}
	fmt.Println("Connected to MongoDB!") // logger

	return client.Database(database), nil
}
