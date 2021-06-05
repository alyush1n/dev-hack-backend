package events

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/user"
	"dev-hack-backend/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(c *gin.Context) {

	username, done := user.ParseBearer(c)
	if done {
		return
	}

	eventsList := make([]model.Event, 0)
	clubsList := make([]model.Club, 0)

	u, _ := db.FindUserByUsername(username)
	for _, clubID := range u.Clubs {
		club := db.GetClubByName(clubID)
		clubsList = append(clubsList, club)
		for _, event := range club.IncomingEvents {
			e, isExist := db.GetEventByID(event)
			if isExist {
				eventsList = append(eventsList, e)
			}

		}
	}
	fmt.Println(eventsList)
	// TODO: sort eventsList by date

	c.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"events_list": eventsList,
		"clubs_list":  clubsList,
	})
}
