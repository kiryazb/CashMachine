package database

import (
	"CashMachine/internal/config"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq" // pq driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDb(cfg *config.Config, logger *slog.Logger) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME)

    db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    logger.Info("Database loaded successfully!", "module", "database", "DB_NAME", cfg.DB_NAME)

	return db
}