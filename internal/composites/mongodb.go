package composites

import (
	"context"
	"dev-hack-backend/pkg/client/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, mongoURI, database string) (*MongoDBComposite, error) {
	log.Println("Creating new mongo client")
	client, err := mongodb.NewClient(ctx, mongoURI, database)
	if err != nil {
		return nil, err
	}

	log.Println("Mongo client complete")
	return &MongoDBComposite{db: client}, nil
}
