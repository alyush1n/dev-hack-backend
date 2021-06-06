package events

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/user"
	"dev-hack-backend/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Create(c *gin.Context) {
	jsonInput := struct {
		Type        string   `json:"type"`
		Clubs       []string `json:"clubs"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Count       int      `json:"count"`
		Location    string   `json:"location"`
		Date        string   `json:"date"`
		URL         string   `json:"url"`
	}{}

	username, done := user.ParseBearer(c)
	if done {
		return
	}
	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	eventID := primitive.NewObjectID()
	attachmentID := primitive.NewObjectID()
	event := model.Event{
		Id:          eventID,
		Type:        jsonInput.Type,
		Name:        jsonInput.Name,
		Description: jsonInput.Description,
		Count:       jsonInput.Count,
		Clubs:       jsonInput.Clubs,
		Location:    jsonInput.Location,
		Date:        jsonInput.Date,
		Attachment: model.Attachment{
			Id:     attachmentID,
			URL:    jsonInput.URL,
			SentBy: username,
		},
	}
	err := db.InsertAttachment(event.Attachment)
	if err != nil {
		fmt.Println("insert attachment: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
	err = db.InsertEvent(event)
	db.AddEventToClub(jsonInput.Clubs[0], eventID)
	if err != nil {
		fmt.Println("insert event: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"event":   event,
	})

}
