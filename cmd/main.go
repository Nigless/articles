package main

import (
	"log"

	"rest-api/config"
	"rest-api/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	// Run
	app := app.New(cfg)
	app.Start()
}
