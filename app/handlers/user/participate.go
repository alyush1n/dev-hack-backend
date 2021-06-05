package user

import (
	"dev-hack-backend/app/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const Distance = 0.5

func Participate(c *gin.Context) {

	jsonInput := struct {
		EventID      string `json:"event_id" bson:"event_id"`
		UserLocation string `json:"user_location" bson:"user_location"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	username, done := ParseBearer(c)
	if done {
		return
	}

	user, ok := db.FindUserByUsername(username)
	if ok {

		xys := strings.Split(jsonInput.UserLocation, " ")
		if len(xys) != 2 {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		x, err := strconv.ParseFloat(xys[0], 64)
		if err != nil {
			fmt.Println(err)
		}
		y, err := strconv.ParseFloat(xys[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		event, ok := db.GetEventByID(jsonInput.EventID)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		eventXYS := strings.Split(event.Location, " ")
		if len(eventXYS) != 2 {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		xEvent, err := strconv.ParseFloat(eventXYS[0], 64)
		if err != nil {
			fmt.Println(err)
		}
		yEvent, err := strconv.ParseFloat(eventXYS[1], 64)
		if err != nil {
			fmt.Println(err)
		}

		if math.Abs(xEvent-x) <= Distance && math.Abs(yEvent-y) <= Distance {

			event.Count++
			if event.Count >= 15 {
				event.AvailablePoints = int(1000 + math.Pow(1.07, float64(event.Count)))
				user.Points += event.AvailablePoints
			}

			db.UpdateEvent(event)
			user.VisitedEvents = append(user.VisitedEvents)
			db.UpdateUser(user)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
			return
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
