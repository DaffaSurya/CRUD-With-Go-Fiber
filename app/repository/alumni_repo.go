package repository

import (
	"crud-app/app/models"
	"database/sql"
)

type AlumniRepository interface {
	FindAll() ([]models.Alumni, error)
	FindByID(id int) (*models.Alumni, error)
	Create(a *models.Alumni) error
	Update(id int, a *models.Alumni) error
	Delete(id int) error
}

type alumniPostgres struct{ db *sql.DB }

func NewAlumniRepository(db *sql.DB) AlumniRepository {
	return &alumniPostgres{db}
}

func (r *alumniPostgres) FindAll() ([]models.Alumni, error) {
	rows, err := r.db.Query("SELECT id, nama, email FROM alumni")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Alumni
	for rows.Next() {
		var a models.Alumni
		rows.Scan(&a.ID, &a.Nama, &a.Email)
		list = append(list, a)
	}
	return list, nil
}

func (r *alumniPostgres) FindByID(id int) (*models.Alumni, error) {
	var a models.Alumni
	err := r.db.QueryRow("SELECT id, nama, email FROM alumni WHERE id=$1", id).
		Scan(&a.ID, &a.Nama, &a.Email)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *alumniPostgres) Create(a *models.Alumni) error {
	_, err := r.db.Exec("INSERT INTO alumni(nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon) VALUES($1,$2,$3,$4,$5,$6,$7)", a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.Tahun_lulus, a.Email, a.No_telp)
	return err
}

func (r *alumniPostgres) Update(id int, a *models.Alumni) error {
	_, err := r.db.Exec("UPDATE alumni SET nama=$1, jurusan=$2, no_telepon=$3 WHERE id=$4", a.Nama, a.Jurusan, a.No_telp, id)
	return err
}

func (r *alumniPostgres) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM alumni WHERE id=$1", id)
	return err
}
