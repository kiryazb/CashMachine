package main

import (
	"CashMachine/internal/api/handlers"
	"CashMachine/internal/config"
	"CashMachine/internal/database"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("main package")

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        AddSource: true,               // Включаем информацию об источнике (файл и номер строки)
        Level:     slog.LevelInfo,     // Устанавливаем уровень логирования INFO
    })

    logger := slog.New(handler)

	cfg := config.MustLoad(logger)

	db := database.ConnectDb(cfg, logger)

	h := &handlers.Handlers{
		DB:     db,
		Logger: logger,
	}

	logger.Info("Setup main app", "module", "main")

	_ = db

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handlers.HelloHandler)

	r.Route("/item", func(r chi.Router) {
		r.Post("/", h.CreateItem)
		r.Get("/{userID}", h.ReadItem)
		r.Put("/{userID}", h.UpdateItem)
		r.Delete("/{userID}", h.DeleteItem)
	})

	logger.Info("Server is listening on port 8080...", "module", "main")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		logger.Error("Server failed to start:", err)
	}

}