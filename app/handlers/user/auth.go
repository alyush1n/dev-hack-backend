package user

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/session"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Auth(c *gin.Context) {

	jsonInput := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}
	//TODO: GetUserByID Andrey
	user, exist := db.FindUserByID(jsonInput.Username)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonInput.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"token":   session.Create(user.Username),
	})

}
