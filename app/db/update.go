package db

import (
	"context"
	"dev-hack-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//func UpdateEvent(event model.Event) (isExist bool) {
//	filter := bson.M{"_id": event.Id}
//
//	update := bson.D{
//		{"$set", bson.D{
//			{"name", event.Name},
//			{"location", event.City},
//			{"date", event.Date},
//			{"attachment", bson.D{
//				{"url", },
//				{"sent_by", event.SentBy},
//			}},
//		}},
//	}
//
//	_, err := eventsCollection.UpdateOne(context.Background(), update, filter)
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return false
//		}
//		return
//	}
//
//	return true
//}

func UpdateUser(User model.User) (isExist bool) {
	filter := bson.M{"_id": User.Id}

	update := bson.D{
		{"$set", bson.D{
			{"username", User.Username},
			{"password", User.Password},
			{"first_name", User.FirstName},
			{"last_name", User.LastName},
			{"sex", User.Sex},
			{"points", User.Points},
		}},
		{"$push", bson.D{
			{"clubs", User.Clubs},
		}},
	}

	_, err := usersCollection.UpdateOne(context.Background(), update, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}
