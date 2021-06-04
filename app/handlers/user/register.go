package user

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"dev-hack-backend/app/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {

	var userToken string

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(jsonInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}
	userID := primitive.NewObjectID()
	user := model.User{
		Id:       userID,
		Username: jsonInput.Username,
		Password: string(hashedPassword),
	}
	err = db.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	userToken, err = session.Create(jsonInput.Username)
	if err != nil {
		fmt.Println("Error in generating JWT: " + err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"token":   userToken,
	})
}
