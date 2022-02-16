package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
	user2 "dev-hack-backend/internal/service/user"
	"dev-hack-backend/pkg/apperror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type userStorage struct {
	database       *mongo.Database
	userCollection string
}

func NewStorage(database *mongo.Database, userCollection string) user2.Storage {
	return &userStorage{
		database:       database,
		userCollection: userCollection,
	}
}

func (s *userStorage) GetUserById(ctx context.Context, id string) (*user.User, error) {
	var currentUser user.User
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": userId}

	err = s.database.Collection(s.userCollection).FindOne(ctx, filter).Decode(&currentUser)
	if err != nil {
		return nil, apperror.MongoFindError
	}

	return &currentUser, nil
}

func (s *userStorage) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	var currentUser user.User
	filter := bson.M{"username": username}

	err := s.database.Collection(s.userCollection).FindOne(ctx, filter).Decode(&currentUser)
	if err != nil {
		return nil, apperror.MongoFindError
	}
	return &currentUser, nil
}

func (s *userStorage) InsertUser(ctx context.Context, currentUser *user.User) error {
	_, err := s.database.Collection(s.userCollection).InsertOne(ctx, currentUser)
	if err != nil {
		return apperror.MongoInsertError
	}

	return nil
}

func (s *userStorage) UpdateUser(ctx context.Context, currentUser *user.User) error {
	filter := bson.M{"_id": currentUser.Id}

	res := s.database.Collection(s.userCollection).FindOneAndReplace(ctx, filter, &currentUser)
	if res.Err() != nil {
		log.Print(res.Err().Error())
		return apperror.MongoUpdateError
	}
	return nil
}

func (s *userStorage) DeleteUserById(ctx context.Context, id string) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userId}

	_, err = s.database.Collection(s.userCollection).DeleteOne(ctx, filter)
	if err != nil {
		return apperror.MongoDeleteError
	}

	return nil
}

func (s *userStorage) UpdateSession(ctx context.Context, id string, session user.Session) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return apperror.ObjectIdError
	}
	filter := bson.M{"_id": userId}

	update := bson.D{
		{"$set", bson.D{
			{"session", session},
		}},
	}

	_, err = s.database.Collection(s.userCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return apperror.MongoUpdateSessionError
	}
	return nil
}

func (s *userStorage) GetUserByRT(ctx context.Context, id string, rToken string) (*user.User, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, apperror.ObjectIdError
	}
	filter := bson.M{"_id": userId}

	var currentUser user.User
	err = s.database.Collection(s.userCollection).FindOne(ctx, filter).Decode(&currentUser)
	if err != nil {
		return nil, apperror.MongoFindError
	}
	if currentUser.Session.RefreshToken != rToken {
		return nil, apperror.RefreshEqualError
	}
	if time.Now().After(currentUser.Session.ExpiresAt) {
		return nil, apperror.RefreshTimeError
	}
	return &currentUser, nil
}
