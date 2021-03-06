package user

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {

	jsonInput := model.User{}
	user, ok := db.FindUserByUsername(jsonInput.Username)
	if ok {

		u := model.User{
			Id:            jsonInput.Id,
			Username:      jsonInput.Username,
			Password:      jsonInput.Password,
			Clubs:         jsonInput.Clubs,
			VisitedEvents: jsonInput.VisitedEvents,
			FirstName:     jsonInput.Username,
			LastName:      jsonInput.LastName,
			Sex:           jsonInput.Sex,
		}

		switch {
		case jsonInput.Sex == "":
			u.Sex = user.Sex
		case jsonInput.LastName == "":
			u.LastName = user.LastName
		case jsonInput.FirstName == "":
			u.FirstName = user.FirstName
		case jsonInput.PhotoURL == "":
			u.PhotoURL = user.PhotoURL
		}

		db.UpdateUser(u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "internal server error",
	})
	return

}

func Me(c *gin.Context) {
	username, done := ParseBearer(c)
	if done {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	user, ok := db.FindUserByUsername(username)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
