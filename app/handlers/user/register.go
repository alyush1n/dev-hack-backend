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
			"message": "internal server errorrrr",
		})
		return
	}
	userID := primitive.NewObjectID()
	user := model.User{
		Id:            userID,
		Type:          "stuff",
		Username:      jsonInput.Username,
		Password:      string(hashedPassword),
		Clubs:         nil,
		VisitedEvents: nil,
		FirstName:     "abc",
		LastName:      "def",
		Sex:           "non-binary",
		Points:        0,
	}

	err = db.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
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
