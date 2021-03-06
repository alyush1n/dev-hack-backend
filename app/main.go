package main

import (
	"dev-hack-backend/app/config"
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/attachments"
	"dev-hack-backend/app/handlers/events"
	"dev-hack-backend/app/handlers/items"
	"dev-hack-backend/app/handlers/user"
	"fmt"
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
	app.POST("/event", events.Create)
	app.GET("/feed", events.Load)
	app.PUT("/user", user.Update)
	app.GET("/user/me", user.Me)
	app.POST("/attachment", attachments.Upload)
	app.POST("/participate", user.Visit)
	app.POST("/registerToEvent", user.RegisterToEvent)
	app.GET("/items",items.Load)
	err := app.Run("localhost:" + config.Port)
	if err != nil {
		fmt.Println("Error in launching backend: " + err.Error())
	}

}
