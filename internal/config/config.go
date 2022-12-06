package configs

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PgSource         string
	MongoDbConnSting string
	ServerPort       string
	Secret           string
	ExpTime          time.Duration
	RedisPassword    string
	RedisHost        string
	GrpcClient       string
}

func LoadConnectionConfig() (*Config, error) {
	var config Config
	err := godotenv.Load("./internal/config/connection.env")
	if err != nil {
		return nil, err
	}

	config.PgSource = os.Getenv("PG_SOURCE")
	config.MongoDbConnSting = os.Getenv("MONGODB_CONNSTRING")
	config.ServerPort = os.Getenv("SERVERPORT")
	config.Secret = os.Getenv("SECRET")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	config.RedisHost = os.Getenv("REDIS_HOST")
	config.GrpcClient = os.Getenv("GRPC_CLIENT")
	config.ExpTime, err = time.ParseDuration(os.Getenv("EXPTIME"))
	if err != nil {
		return nil, err
	}

	return &config, err
}
