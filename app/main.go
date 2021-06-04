package main

import (
	"dev-hack-backend/app/config"
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/user"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	db.Connect()

	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "not found"})
	})

	app.POST("/auth", user.Auth)
	app.POST("/user", user.Register)
}
