package main

import (
	"log"

	"github.com/AnxVit/ozon_1/internal/core/web"
)

func main() {
	err := web.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize the app: %v", err)
	}
}
