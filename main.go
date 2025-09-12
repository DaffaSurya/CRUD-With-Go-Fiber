package main

import (
	"crud-app/database"
	"crud-app/routes"

	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "12345678")
	os.Setenv("DB_NAME", "mahasiswa_db")

	database.ConnectDB()

	app := fiber.New()

	// alumni 
	// routes.SetupRoutes(app)
	routes.SetupRoutes(app)
	app.Listen(":3000")
}