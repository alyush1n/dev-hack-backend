package db

import (
	"context"
	"dev-hack-backend/app/model"
)

func InsertUser(User model.User) (err error) {
	_, err = usersCollection.InsertOne(context.Background(), User)
	if err != nil {
		return err
	}
	return nil
}

func InsertEvent(Event model.Event) (isExist bool) {
	_, err := eventCollection.InsertOne(context.Background(), Event)
	if err != nil {
		return false
	}
	return true
}
