package main

import (
	handlers "crud-app/Handlers"
	middleware "crud-app/Middleware"
	"crud-app/app/repository"
	"crud-app/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	db := database.ConnectDB()

	// repositories
	userRepo := repository.NewUserRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	pekerjaanRepo := repository.NewPekerjaanRepository(db)

	// handlers
	authHandler := handlers.NewAuthHandler(userRepo)
	alumniHandler := handlers.NewAlumniHandler(alumniRepo)
	pekerjaanHandler := handlers.NewPekerjaanHandler(pekerjaanRepo)

	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Alumni routes
	r.Handle("/alumni", middleware.AuthMiddleware(userRepo, http.HandlerFunc(alumniHandler.GetAll))).Methods("GET")
	r.Handle("/alumni/{id}", middleware.AuthMiddleware(userRepo, http.HandlerFunc(alumniHandler.GetByID))).Methods("GET")

	r.Handle("/alumni", middleware.AuthMiddleware(userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(alumniHandler.Create)))).Methods("POST")

	r.Handle("/alumni/{id}", middleware.AuthMiddleware(userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(alumniHandler.Update)))).Methods("PUT")

	r.Handle("/alumni/{id}", middleware.AuthMiddleware(userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(alumniHandler.Delete)))).Methods("DELETE")

	// Pekerjaan routes
	r.Handle("/pekerjaan", middleware.AuthMiddleware(userRepo, http.HandlerFunc(pekerjaanHandler.GetAll))).Methods("GET")
	r.Handle("/pekerjaan/{alumni_id}", middleware.AuthMiddleware(userRepo, http.HandlerFunc(pekerjaanHandler.GetByAlumni))).Methods("GET")
	r.Handle("/pekerjaan", middleware.AuthMiddleware(userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(pekerjaanHandler.Create)))).Methods("POST")
	r.Handle("/pekerjaan/{id}", middleware.AuthMiddleware(userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(pekerjaanHandler.Update)))).Methods("PUT")
	r.Handle("/pekerjaan/{id}", middleware.AuthMiddleware(userRepo,
		middleware.RoleMiddleware("admin", http.HandlerFunc(pekerjaanHandler.Delete)))).Methods("DELETE")

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("🚀 Server running at http://localhost:" + port)
	http.ListenAndServe(":"+port, r)
}
