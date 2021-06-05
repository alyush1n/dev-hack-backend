package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	AccessSecret string
	MongoPass    string
	Port         string
	S3ID         string
	S3Secret     string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading configs: " + err.Error())
	} else {
		AccessSecret = os.Getenv("ACCESS_SECRET")
		MongoPass = os.Getenv("MONGO_PASS")
		Port = os.Getenv("PORT")
		S3ID = os.Getenv("S3_ID")
		S3Secret = os.Getenv("S3_SECRET")
	}
}
