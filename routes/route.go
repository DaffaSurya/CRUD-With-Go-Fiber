package routes

import (
	middleware "crud-app/Middleware"
	service "crud-app/app/Service"
	"crud-app/app/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router, PekerjaanService *service.PekerjaanService, alumniService *service.AlumniService, authService *service.AuthService, userRepo *repository.UserRepository) {
	r.HandleFunc("/register", authService.Register).Methods("POST")
	r.HandleFunc("/login", authService.Login).Methods("POST")

	// Alumni routes
	r.Handle("/alumni", middleware.AuthMiddleware(*userRepo, http.HandlerFunc(alumniService.GetAll))).Methods("GET")
	r.Handle("/alumni/{id}", middleware.AuthMiddleware(*userRepo, http.HandlerFunc(alumniService.GetByID))).Methods("GET")

	r.Handle("/alumni", middleware.AuthMiddleware(*userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(alumniService.Create)))).Methods("POST")

	r.Handle("/alumni/{id}", middleware.AuthMiddleware(*userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(alumniService.Update)))).Methods("PUT")

	r.Handle("/alumni/{id}", middleware.AuthMiddleware(*userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(alumniService.Delete)))).Methods("DELETE")

	// Pekerjaan routes
	r.Handle("/pekerjaan", middleware.AuthMiddleware(*userRepo, http.HandlerFunc(PekerjaanService.GetAll))).Methods("GET")
	r.Handle("/pekerjaan/{alumni_id}", middleware.AuthMiddleware(*userRepo, http.HandlerFunc(PekerjaanService.GetByAlumni))).Methods("GET")
	r.Handle("/pekerjaan", middleware.AuthMiddleware(*userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(PekerjaanService.Create)))).Methods("POST")
	r.Handle("/pekerjaan/{id}", middleware.AuthMiddleware(*userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(PekerjaanService.Update)))).Methods("PUT")
	r.Handle("/pekerjaan/{id}", middleware.AuthMiddleware(*userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(PekerjaanService.Delete)))).Methods("DELETE")
}