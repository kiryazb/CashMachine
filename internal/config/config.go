package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST string
	DB_PORT int
	DB_NAME string
	DB_USER string
	DB_PASS string
}

func MustLoad() *Config {
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	cfg := &Config{}

	cfg.DB_HOST = os.Getenv("DB_HOST")
	cfg.DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
    if err != nil {
        log.Fatalf("Error converting DB_PORT to integer: %v", err)
    }
	cfg.DB_NAME = os.Getenv("DB_NAME")
	cfg.DB_USER = os.Getenv("DB_USER")
	cfg.DB_PASS = os.Getenv("DB_PASS")

	
	return cfg
}

