package service


import (
	"strconv"
	"time"

	"crud-app/app/models"
	"crud-app/app/repository"

	"github.com/gofiber/fiber/v2"
)

// type PekerjaanAlumniHandler struct {
// 	repo repository.PekerjaanAlumniRepository
// }

type PekerjaanAlumniHandler struct {
	repo repository.PekerjaanAlumniRepository
}

func NewPekerjaanAlumniHandler(r repository.PekerjaanAlumniRepository) *PekerjaanAlumniHandler {
	return &PekerjaanAlumniHandler{repo: r}
}

// CREATE
func (h *PekerjaanAlumniHandler) Create(c *fiber.Ctx) error {
	var job models.Pekerjaan_Alm
	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

	if err := h.repo.Create(&job); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(job)
}


// GET by AlumniID
func (h *PekerjaanAlumniHandler) GetByAlumniID(c *fiber.Ctx) error {
	alumniID, _ := strconv.Atoi(c.Params("alumni_id"))
	jobs, err := h.repo.FindByAlumniID(alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(jobs)
}

// UPDATE
func (h *PekerjaanAlumniHandler) Update(c *fiber.Ctx) error {
	var job models.Pekerjaan_Alm
	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	job.UpdatedAt = time.Now()

	if err := h.repo.Update(&job); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(job)
}

// DELETE
func (h *PekerjaanAlumniHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}






