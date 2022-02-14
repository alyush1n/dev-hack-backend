package api

import (
	"dev-hack-backend/internal/domain/user"
	"github.com/gin-gonic/gin"
)

func NewResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func ResponseWithTokens(c *gin.Context, code int, aToken, rToken string) {
	c.JSON(code, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

func ResponseUser(c *gin.Context, code int, user *user.User) {
	c.JSON(code, gin.H{
		"username":       user.Username,
		"password":       user.Password,
		"first_name":     user.FirstName,
		"last_name":      user.LastName,
		"clubs":          user.Clubs,
		"visited_events": user.VisitedEvents,
		"photo_url":      user.PhotoURL,
		"sex":            user.Sex,
		"points":         user.Points,
	})
}
