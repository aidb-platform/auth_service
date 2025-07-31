package main

import (
	"github.com/aidb-platform/auth_service/config"
	"github.com/aidb-platform/auth_service/middleware"
	"github.com/aidb-platform/auth_service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDatabase()

	r := gin.Default()

	r.POST("/signup", routes.SignUp(db))
	r.POST("/login", routes.Login(db))

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/me", routes.CurrentUser(db))

	r.Run(":8080")
}
