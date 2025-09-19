package handlers

import (
	"crud-app/app/models"
	"crud-app/app/repository"
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type PekerjaanHandler struct {
	repo repository.PekerjaanRepository
}

func NewPekerjaanHandler(r repository.PekerjaanRepository) *PekerjaanHandler {
	return &PekerjaanHandler{repo: r}
}

func (h *PekerjaanHandler) GetByAlumni(w http.ResponseWriter, r *http.Request) {
	alumniID, _ := strconv.Atoi(mux.Vars(r)["alumni_id"])
	data, err := h.repo.FindByAlumni(alumniID)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *PekerjaanHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Pekerjaan
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&p); err != nil {
		http.Error(w, "Failed to create pekerjaan", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *PekerjaanHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var p models.Pekerjaan
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(id, &p); err != nil {
		http.Error(w, "Failed to update pekerjaan", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Updated successfully"})
}

func (h *PekerjaanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, "Failed to delete pekerjaan", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}

func (h *PekerjaanHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pekerjaan, err := h.repo.FindAllPekerjaan()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pekerjaan)
}