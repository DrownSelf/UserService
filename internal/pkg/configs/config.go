package configs

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	PgSource         string
	MongoDbConnSting string
	ServerPort       string
	Secret           string
	ExpTime          time.Duration
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
	config.ExpTime, err = time.ParseDuration(os.Getenv("EXPTIME"))
	if err != nil {
		return nil, err
	}

	return &config, err
}
