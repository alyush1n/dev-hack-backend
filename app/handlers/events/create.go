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
		Name            string           `json:"name" bson:"name" binding:"required"`
		City            string           `json:"city" bson:"city" binding:"required"`
		Date            string           `json:"date" bson:"date"`
		Logo            model.Attachment `json:"logo" bson:"logo"`
		BackgroundImage model.Attachment `json:"background_image" bson:"background_image"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	eventID := primitive.NewObjectID()
	event := model.Event{
		Id:              eventID,
		Name:            jsonInput.Name,
		City:            jsonInput.City,
		Date:            jsonInput.Date,
		Logo:            jsonInput.Logo,
		BackgroundImage: jsonInput.BackgroundImage,
	}

	err := db.InsertEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
}
