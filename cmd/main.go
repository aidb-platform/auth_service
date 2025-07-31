package main

import (
	"log"

	"aidb-auth-service/config"
	"aidb-auth-service/models"
	"aidb-auth-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	models.ConnectDatabase(cfg.DatabaseURL)
	models.DB.AutoMigrate(&models.User{}, &models.Organization{})

	r := gin.Default()
	routes.Setup(r)

	log.Printf("Starting server on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
