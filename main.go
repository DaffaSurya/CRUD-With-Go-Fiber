package main

import (
	"crud-app/database"
	"crud-app/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "12345678")
	os.Setenv("DB_NAME", "mahasiswa_db")

	database.ConnectDB()

	app := fiber.New()

	app.Post("/Mahasiswa", handlers.CreateUser)
	app.Get("/Mahasiswa", handlers.GetUsers)
	app.Put("/Mahasiswa/:id", handlers.UpdateUser)
	app.Delete("/Mahasiswa/:id", handlers.DeleteUser)

	app.Listen(":3000")
}