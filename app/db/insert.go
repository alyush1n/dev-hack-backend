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

func InsertEvent(Event model.Event) (err error) {
	_, err = eventsCollection.InsertOne(context.Background(), Event)
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
