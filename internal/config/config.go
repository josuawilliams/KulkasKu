package config

import (
	"fmt"
	"kulkasku/internal/model"
	"os"

	"github.com/joho/godotenv"
)


func LoadConfigDatabase() (*model.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed Load Env")
	}

	return &model.Config{
		Port:        os.Getenv("APP_PORT"),
		DBUrlMigrate:os.Getenv("DATABASE_URL"),
		SecretJwt:   os.Getenv("SECRET_JWT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
	}, nil


}
