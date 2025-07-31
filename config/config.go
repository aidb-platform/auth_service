package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EnvConfig struct {
	DatabaseURL string
	JWTSecret   string
}

var Env EnvConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	Env = EnvConfig{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(Env.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("✅ Connected to DB")
	return db
}
