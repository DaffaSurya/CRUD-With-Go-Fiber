package handlers

import (
	"crud-app/app/models"
	"crud-app/app/repository"
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type AlumniHandler struct {
	repo repository.AlumniRepository
}

func NewAlumniHandler(r repository.AlumniRepository) *AlumniHandler {
	return &AlumniHandler{repo: r}
}

func (h *AlumniHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	alumni, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(alumni)
}

func (h *AlumniHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	alumni, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, "Alumni not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(alumni)
}

func (h *AlumniHandler) Create(w http.ResponseWriter, r *http.Request) {
	var a models.Alumni
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&a); err != nil {
		http.Error(w, "Failed to create alumni", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(a)
}

func (h *AlumniHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var a models.Alumni
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(id, &a); err != nil {
		http.Error(w, "Failed to update alumni", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Updated successfully"})
}

func (h *AlumniHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Failed to delete alumni", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}
