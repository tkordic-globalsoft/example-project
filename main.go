package main

import (
	"example/core/application/service"
	"example/infrastructure/postgres"
	"example/infrastructure/repository"
	"example/presentation/rest"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConfig := &postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := postgres.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Adapters
	postRepository := repository.NewPostSqlAdapter(db)

	// Services
	postService := service.NewPostService(postRepository)

	// Router
	router := rest.NewRouter(postService)

	router.Run()
}
