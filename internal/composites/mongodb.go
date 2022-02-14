package composites

import (
	"context"
	"dev-hack-backend/pkg/client/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, mongoURI, database string) (*MongoDBComposite, error) {
	client, err := mongodb.NewClient(ctx, mongoURI, database)
	if err != nil {
		return nil, err
	}

	return &MongoDBComposite{db: client}, nil
}
