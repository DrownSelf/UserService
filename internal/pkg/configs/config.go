package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type SecretConfig struct {
	SECRET string
}

type ConnectionConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMODE  string
}

func makeConnectionString(config ConnectionConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.HOST, config.PORT, config.USER, config.PASSWORD, config.DBNAME, config.SSLMODE)
}

func LoadConnectionConfig() (string, error) {
	var config ConnectionConfig
	viper.AddConfigPath("internal/pkg/configs/")
	viper.SetConfigName("connection")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return "", err
	}
	return makeConnectionString(config), err
}

func LoadSecretConfig() (string, error) {
	var secret SecretConfig
	viper.AddConfigPath("internal/pkg/configs/")
	viper.SetConfigName("secret")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	err = viper.Unmarshal(&secret)
	if err != nil {
		return "", err
	}
	return secret.SECRET, err
}
