package events

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Create(c *gin.Context) {

	jsonInput := struct {
		Type             string `json:"type"`
		Name             string `json:"name"`
		Location         string `json:"location"`
		Date             string `json:"date"`
		model.Attachment `json:"attachment"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	eventID := primitive.NewObjectID()
	event := model.Event{
		Id:         eventID,
		Type:       jsonInput.Type,
		Name:       jsonInput.Name,
		Location:   jsonInput.Location,
		Date:       jsonInput.Date,
		Attachment: model.Attachment{},
	}

	db.InsertEvent(event)

}
