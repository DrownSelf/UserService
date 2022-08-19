package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PgSource         string
	MongoDbConnSting string
	ServerPort       string
	Secret           string
	Reset            bool
}

func LoadConnectionConfig() (*Config, error) {
	var config Config
	err := godotenv.Load("./internal/pkg/configs/connection.env")
	if err != nil {
		return nil, err
	}
	config.PgSource = os.Getenv("PG_SOURCE")
	config.MongoDbConnSting = os.Getenv("MONGODB_CONNSTRING")
	config.ServerPort = os.Getenv("SERVERPORT")
	config.Secret = os.Getenv("SECRET")
	reset := os.Getenv("RESET")

	if reset != "true" {
		config.Reset = false
	} else {
		config.Reset = true
	}
	return &config, err
}
