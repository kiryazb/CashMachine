package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST string
	DB_PORT string
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

	cfg.DB_NAME = os.Getenv("DB_NAME")

	fmt.Println(cfg.DB_NAME)

	
	return nil
}

