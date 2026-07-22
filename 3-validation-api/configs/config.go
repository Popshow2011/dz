package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type VerifyConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadCofig() *VerifyConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &VerifyConfig{
		Email:    os.Getenv("Email"),
		Password: os.Getenv("Password"),
		Address:  os.Getenv("Address"),
	}
}
