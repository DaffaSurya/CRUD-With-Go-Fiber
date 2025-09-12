package service

import (
	"strconv"
	"time"

	"crud-app/app/models"
	"crud-app/app/repository"

	"github.com/gofiber/fiber/v2"
)

type AlumniHandler struct {
	repo repository.AlumniRepository
}

func NewAlumniHandler(r repository.AlumniRepository) *AlumniHandler {
	return &AlumniHandler{repo: r}
}

// CREATE
func (h *AlumniHandler) Create(c *fiber.Ctx) error {
	var alumni models.Alumni
	if err := c.BodyParser(&alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	alumni.CreatedAt = time.Now()
	alumni.UpdatedAt = time.Now()

	if err := h.repo.Create(&alumni); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(alumni)
}

// GET ALL
func (h *AlumniHandler) GetAll(c *fiber.Ctx) error {
	alumnis, err := h.repo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumnis)
}


func (h *AlumniHandler) GetID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	alumnis, err := h.repo.FindByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumnis)
}

// UPDATE
func (h *AlumniHandler) Update(c *fiber.Ctx) error {
	var alumni models.Alumni
	if err := c.BodyParser(&alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	alumni.UpdatedAt = time.Now()

	if err := h.repo.Update(&alumni); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(alumni)
}

// DELETE
func (h *AlumniHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
