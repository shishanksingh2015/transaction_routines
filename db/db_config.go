package db

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBHost = os.Getenv("DB_HOST")
	config.ServerAddress = os.Getenv("SERVER_ADDRESS")

	return config, err
}
