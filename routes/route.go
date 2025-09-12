package routes

import (
	"crud-app/app/repository"
	"crud-app/app/service"
	"crud-app/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Repository
	alumniRepo := repository.NewAlumniRepository(database.DB)
	// JobRepo := repository.NewAlumniRepository(database.DB)
	
	// service
	alumniservice := service.NewAlumniHandler(alumniRepo)
    // pekerjaanservice := service.NewPekerjaanAlumniHandler(JobRepo)


	
	// Alumni Routes
	app.Post("/alumni", alumniservice.Create)
	app.Get("/alumni", alumniservice.GetAll)
	app.Get("/alumni/:id", alumniservice.GetID)
	app.Put("/alumni/:id", alumniservice.Update)
	app.Delete("/alumni/:id", alumniservice.Delete)

	// Pekerjaan Alumni Routes
	// api.Post("/pekerjaanAlumni", pekerjaanservice.Create.Create)
	// api.Get("/pekerjaanAlumni/:alumni_id", pekerjaanhandlers.GetByAlumniID)
	// api.Put("/pekerjaanAlumni", pekerjaanhandlers.Update)
	// api.Delete("/pekerjaanAlumni/:id", pekerjaanhandlers.Delete)
	// kalau ada, bisa tambahkan Update / Delete juga
}
