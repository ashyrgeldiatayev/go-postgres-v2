package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/ashyrgeldiatayev/go-postgres/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := internal.LoadConfig()

	db, err := internal.NewConnection(config)
	if err != nil {
		log.Fatal("Could not connect to the database")
	}
	defer db.Close(context.Background())

	err = internal.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate database")
	}

	repo := internal.NewRepository(db)
	controller := internal.NewController(repo)

	app := fiber.New()
	controller.SetupRoutes(app)
	app.Listen(":8080")
}
