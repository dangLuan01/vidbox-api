package main

import (
	"log"

	"vidbox-api/internal/app"
	"vidbox-api/internal/config"
	"vidbox-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	app.LoadEnv()
	
	mode := utils.GetEnv("GIN_MODE", "release")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)	
	}
	
	cfg := config.NewConfig()
	
	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Error start app:%s", err)
	}
	
	if err := application.Run(); err != nil {
		log.Fatalf("Error run app:%s", err)
	}
}