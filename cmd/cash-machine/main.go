package main

import (
	"CashMachine/internal/config"
	"CashMachine/internal/database"
	"fmt"
)

func main() {
	fmt.Println("main package")

	cfg := config.MustLoad()

	db := database.ConnectDb(cfg)

	fmt.Println(db)
}