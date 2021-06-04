package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading configs: " + err.Error())
	} else {

	}
}
