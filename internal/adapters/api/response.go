package user

import (
	"dev-hack-backend/internal/domain/user"
	"github.com/gin-gonic/gin"
)

func NewResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func NewResponseStatus(c *gin.Context, code int) {
	c.Status(code)
}

func NewAbortResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}

func ResponseWithTokens(c *gin.Context, code int, userID, aToken, rToken string) {
	c.JSON(code, gin.H{
		"id":            userID,
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

func ResponseUser(c *gin.Context, code int, user *user.User) {
	c.JSON(code, gin.H{
		"username":       user.Username,
		"first_name":     user.FirstName,
		"last_name":      user.LastName,
		"clubs":          user.Clubs,
		"visited_events": user.VisitedEvents,
		"photo_url":      user.PhotoURL,
		"sex":            user.Sex,
		"points":         user.Points,
	})
}
