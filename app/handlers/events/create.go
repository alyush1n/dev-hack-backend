package events

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Create(c *gin.Context) {
	jsonInput := struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Count    int    `json:"count"`
		Location string `json:"location"`
		Date     string `json:"date"`
		SentBy   string `json:"sent_by"`
		URL      string `json:"url"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	eventID := primitive.NewObjectID()
	event := model.Event{
		Id:       eventID,
		Type:     jsonInput.Type,
		Name:     jsonInput.Name,
		Count:    jsonInput.Count,
		Location: jsonInput.Location,
		Date:     jsonInput.Date,
		Attachment: model.Attachment{
			Id:     primitive.NewObjectID(),
			URL:    jsonInput.URL,
			SentBy: jsonInput.SentBy,
		},
	}
	err := db.InsertAttachment(event.Attachment)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
	err = db.InsertEvent(event)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"event":   event,
	})

}
