package main

import (
	"context"
	"dev-hack-backend/internal/composites"
	"dev-hack-backend/internal/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()
	gin.SetMode(cfg.App.Mode)

	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.MongoURI, cfg.MongoDB.MongoDatabase)
	if err != nil {
		log.Fatal(err)
	}
	userComposite := composites.NewUserComposite(mongoDBC, cfg.App.JWTSecret, cfg.MongoDB.MongoUserCollection, cfg.App.JWTAccessTTL, cfg.App.JWTRefreshTTL)

	userComposite.Handler.Register(router)

	err = router.Run(cfg.Listen.IP + ":" + cfg.Listen.Port)
	if err != nil {
		log.Fatal(err)
	}
}
