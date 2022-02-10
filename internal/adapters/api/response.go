package api

import "github.com/gin-gonic/gin"

func NewResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func ResponseWithTokens(c *gin.Context, code int, message, aToken, rToken string) {
	c.JSON(code, gin.H{
		"message":       message,
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
