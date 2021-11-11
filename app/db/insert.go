package db

import (
	"context"
	"dev-hack-backend/app/model"
)

func InsertUser(user model.User) (err error) {
	_, err = usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func InsertEvent(event model.Event) (err error) {
	_, err = eventsCollection.InsertOne(context.Background(), event)
	if err != nil {
		return err
	}
	return nil
}

func InsertClub(club model.Club) (err error) {
	_, err = clubsCollection.InsertOne(context.Background(), club)
	if err != nil {
		return err
	}
	return nil
}

func InsertAttachment(a model.Attachment) (err error) {
	_, err = attachmentCollection.InsertOne(context.Background(), a)
	if err != nil {
		return err
	}
	return nil
}
