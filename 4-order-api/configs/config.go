package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Db   DbConfig
}

type DbConfig struct {
	DSN         string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_NAME     string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Don`t loaded env file")
	}

	return &Config{
		Port: os.Getenv("PORT"),
		Db: DbConfig{
			DSN:         os.Getenv("DSN"),
			DB_PORT:     os.Getenv("DB_PORT"),
			DB_USERNAME: os.Getenv("DB_USERNAME"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_NAME:     os.Getenv("DB_NAME"),
		},
	}
}
