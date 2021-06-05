package events

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/user"
	"dev-hack-backend/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(c *gin.Context) {

	username, done := user.ParseBearer(c)
	if done {
		return
	}

	list := make([]model.Event, 0)

	u, _ := db.FindUserByUsername(username)
	for _, clubID := range u.Clubs {
		club := db.GetClubByID(clubID)
		for _, eventID := range club.IncomingEvents {
			event, _ := db.GetEventByID(eventID)
			list = append(list, event)
		}

	}
	// TODO: sort list by date

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"events":  list,
	})
}
