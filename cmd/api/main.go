package main

import (
	"log"

	"vidbox-api/internal/app"
	"vidbox-api/internal/config"
)

func main() {
	
	app.LoadEnv()
	
	cfg := config.NewConfig()
	
	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Error start app:%s", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Error run app:%s", err)
	}
}