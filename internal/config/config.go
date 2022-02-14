package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
	"time"
)

type Config struct {
	Listen struct {
		IP   string `yaml:"ip" env-default:"localhost"`
		Port string `yaml:"port" env-default:"8080"`
	}
	MongoDB struct {
		MongoURI            string `yaml:"mongoURI" env-required:"true"`
		MongoDatabase       string `yaml:"mongo_database" env-required:"true"`
		MongoUserCollection string `yaml:"mongo_user_collection" env-required:"true"`
	}
	App struct {
		Mode          string        `yaml:"mode" env-default:"debug"`
		JWTSecret     string        `yaml:"jwt_secret" env-required:"true"`
		JWTAccessTTL  time.Duration `yaml:"jwt_access_ttl" env-required:"true"`
		JWTRefreshTTL time.Duration `yaml:"jwt_refresh_ttl" env-required:"true"`
	}
}

var once sync.Once

func GetConfig() *Config {
	var instance *Config
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
