package handlers

import (
	"crud-app/database"
	"crud-app/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CREATE (POST)

func CreateUser(c *fiber.Ctx) error {
	var Mahasiswa models.Mahasiswa
	if err := c.BodyParser(&Mahasiswa); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	query := "INSERT INTO mahasiswa (nim, nama , jurusan, angkatan, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := database.DB.QueryRow(query, Mahasiswa.NIM, Mahasiswa.Nama, Mahasiswa.Jurusan, Mahasiswa.Angkatan, Mahasiswa.Email, time.Now(),time.Now()).Scan(&Mahasiswa.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(Mahasiswa)
}

// READ ALL (GET)
func GetUsers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at  FROM mahasiswa ORDER BY created_at DESC")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	users := []models.Mahasiswa{}
	for rows.Next() {
		var Mahasiswa models.Mahasiswa
		if err := rows.Scan(&Mahasiswa.ID, &Mahasiswa.NIM, &Mahasiswa.Nama, &Mahasiswa.Jurusan, &Mahasiswa.Angkatan, &Mahasiswa.Email, &Mahasiswa.CreatedAt, &Mahasiswa.UpdatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		users = append(users, Mahasiswa)
	}

	return c.JSON(users)
}

// UPDATE (PUT)
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var Mahasiswa models.Mahasiswa
	if err := c.BodyParser(&Mahasiswa); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	query := "UPDATE mahasiswa SET nama = $1, jurusan = $2, angkatan = $3, email = $4, updated_at = $5 WHERE id = $6"
	_, err := database.DB.Exec(query, Mahasiswa.Nama, Mahasiswa.Jurusan, Mahasiswa.Angkatan, Mahasiswa.Email, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	Mahasiswa.ID = atoi(id)
	return c.JSON(Mahasiswa)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM mahasiswa WHERE id = $1", id)
	if err != nil {
		// return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus mahasiswa",
		})
	}
	return c.SendStatus(204)
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
