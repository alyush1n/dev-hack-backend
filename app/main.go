package main

import (
	"dev-hack-backend/app/config"
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/events"
	"dev-hack-backend/app/handlers/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()
	fmt.Println(config.MongoPass)
	db.Connect()

	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "not found"})
	})
	app.POST("/auth", user.Auth)
	app.POST("/user", user.Register)
	app.POST("/event",events.Create)
	app.GET("/feed", events.Load)

	err := app.Run("localhost:" + config.Port)
	if err != nil {
		fmt.Println("Error in launching backend: " + err.Error())
	}
}
