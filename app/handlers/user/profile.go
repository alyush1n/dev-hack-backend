package user

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {

	jsonInput := model.User{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	isExist := db.UpdateUser(jsonInput)
	if !isExist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	return

}
