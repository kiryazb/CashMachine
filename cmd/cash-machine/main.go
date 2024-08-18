package main

import (
	"config/config"
	"fmt"
)

func main() {
	fmt.Println("main package")
	config.MustLoad()
}