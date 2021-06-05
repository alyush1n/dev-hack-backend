package events

import (
	"dev-hack-backend/app/config"
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/auth/pkg/auth"
	"github.com/zhashkevych/auth/pkg/parser"
	"net/http"
	"strings"
)

func Load(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	username, err := parser.ParseToken(headerParts[1], []byte(config.AccessSecret))
	if err != nil {
		status := http.StatusBadRequest
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	var list []model.Event

	user, _ := db.FindUserByUsername(username)
	for _, clubID := range user.Clubs {
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
