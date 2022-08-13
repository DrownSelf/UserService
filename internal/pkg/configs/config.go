package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host        string `mapstructure:"PG_HOST"`
	PostgrePort string `mapstructure:"PG_PORT"`
	User        string `mapstructure:"PG_USER"`
	Password    string `mapstructure:"PG_PASSWORD"`
	Dbname      string `mapstructure:"PG_DB"`
	Sslmode     string `mapstructure:"PG_SSLMODE"`
	ServerPort  string `mapstructure:"SERVERPORT"`
	Secret      string `mapstructure:"SECRET"`
}

func MakeConnectionString(config Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.PostgrePort, config.User, config.Password, config.Dbname, config.Sslmode)
}

func LoadConnectionConfig() (*Config, error) {
	var config Config
	viper.AddConfigPath("./internal/pkg/configs/")
	viper.SetConfigName("connection")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
