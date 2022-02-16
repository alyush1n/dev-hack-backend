package main

import (
	"context"
	"dev-hack-backend/internal/composites"
	"dev-hack-backend/internal/config"
	"dev-hack-backend/pkg/shutdown"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {
	log.Println("init config")
	cfg := config.GetConfig()

	log.Println("init router")
	router := gin.Default()
	gin.SetMode(cfg.App.Mode)
	server := &http.Server{
		Addr:           cfg.Listen.IP + ":" + cfg.Listen.Port,
		Handler:        router,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("init mongo composite")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.MongoURI, cfg.MongoDB.MongoDatabase)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("init user composite")
	userComposite := composites.NewUserComposite(mongoDBC, cfg.App.JWTSecret, cfg.MongoDB.MongoUserCollection, cfg.App.JWTAccessTTL, cfg.App.JWTRefreshTTL)

	userComposite.Handler.Register(router)

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, server)

	log.Println("start serve")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
