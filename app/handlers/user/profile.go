package user

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {

	jsonInput := struct {
		Username  string `json:"username" bson:"username"`
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
		Sex       string `json:"sex" bson:"sex"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}
	user, ok := db.FindUserByUsername(jsonInput.Username)
	if ok {
		u := model.User{
			Id:            user.Id,
			Username:      user.Username,
			Password:      user.Password,
			Clubs:         user.Clubs,
			VisitedEvents: user.VisitedEvents,
			FirstName:     jsonInput.Username,
			LastName:      jsonInput.LastName,
			Sex:           jsonInput.Sex,
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
