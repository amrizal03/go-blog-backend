package main

import (
	"blogbackend/database"
	"blogbackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000",
	}))
}
