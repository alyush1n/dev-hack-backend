package user

import (
	"context"
	"dev-hack-backend/internal/domain/user"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	mongoError        = "failed to %s user with err %w"
	objectIdError     = "error with convert string to ObjectId"
	refreshEqualError = "refresh token is invalid"
	refreshTimeError  = "failed to refresh token (time is up)"
)

type userStorage struct {
	database       *mongo.Database
	userCollection string
}

func NewStorage(database *mongo.Database, userCollection string) user.Storage {
	return &userStorage{
		database:       database,
		userCollection: userCollection,
	}
}

func (s *userStorage) GetUserById(ctx context.Context, id string) (*user.User, error) {
	var user user.User
	userId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": userId}

	err = s.database.Collection(s.userCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf(mongoError, "get", err)
	}
	return &user, nil
}

func (s *userStorage) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	var user user.User
	filter := bson.M{"username": username}

	err := s.database.Collection(s.userCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf(mongoError, "get", err)
	}
	return &user, nil
}

func (s *userStorage) InsertUser(ctx context.Context, user *user.User) (*user.User, error) {
	_, err := s.database.Collection(s.userCollection).InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf(mongoError, "insert", err)
	}
	return user, nil
}

func (s *userStorage) UpdateUser(ctx context.Context, user *user.User) (*user.User, error) {
	filter := bson.M{"_id": user.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", user.Username},
			{"password", user.Password},
			{"photo_url", user.PhotoURL},
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"sex", user.Sex},
			{"points", user.Points},
			{"session", user.Session},
		}},
		{"$push", bson.D{
			{"clubs", user.Clubs},
		}},
	}

	_, err := s.database.Collection(s.userCollection).UpdateOne(ctx, update, filter)
	if err != nil {
		return nil, fmt.Errorf(mongoError, "update", err)
	}
	return user, nil
}

func (s *userStorage) DeleteUserById(ctx context.Context, id string) error {
	userId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": userId}

	_, err = s.database.Collection(s.userCollection).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf(mongoError, "delete", err)
	}

	return nil
}

func (s *userStorage) UpdateSession(ctx context.Context, id string, session user.Session) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf(objectIdError)
	}
	filter := bson.M{"_id": userId}

	update := bson.D{
		{"$set", bson.D{
			{"session", session},
		}},
	}

	_, err = s.database.Collection(s.userCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf(mongoError, "update session", err)
	}
	return nil
}

func (s *userStorage) GetUserByRT(ctx context.Context, id string, rToken string) (*user.User, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf(objectIdError)
	}
	filter := bson.M{"_id": userId}

	var user user.User
	err = s.database.Collection(s.userCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf(mongoError, "get", err)
	}
	if user.Session.RefreshToken != rToken {
		return nil, fmt.Errorf(refreshEqualError)
	}
	if time.Now().After(user.Session.ExpiresAt) {
		return nil, fmt.Errorf(refreshTimeError)
	}
	return &user, nil
}
