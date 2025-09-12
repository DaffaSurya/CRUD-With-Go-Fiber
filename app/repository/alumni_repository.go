package repository

import (
	"crud-app/app/models"
	"database/sql"
	"time"
)

type AlumniRepository interface {
	Create(alumni *models.Alumni) error
	FindAll() ([]models.Alumni, error)
	FindByID(id int) ([]models.Alumni, error)
	Update(alumni *models.Alumni) error
	Delete(id int) error
}

type alumniRepo struct {
	db *sql.DB
}

func NewAlumniRepository(db *sql.DB) AlumniRepository {
	return &alumniRepo{db}
}

func (r *alumniRepo) Create(alumni *models.Alumni) error {
	query := "INSERT INTO alumni (nim, nama , jurusan, angkatan,  tahun_lulus, email, no_telepon, alamat , created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	return r.db.QueryRow(query,
	alumni.NIM, alumni.Nama, alumni.Jurusan, alumni.Angkatan, alumni.Tahun_lulus, alumni.Email, alumni.No_telp, alumni.Alamat, time.Now(),time.Now()).Scan(&alumni.ID)
	
}

func (r *alumniRepo) FindAll() ([]models.Alumni, error) {
	rows, err := r.db.Query("SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, created_at, updated_at up  FROM alumni ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Alumni
	for rows.Next() {
		var Alumni models.Alumni
		if err := rows.Scan(&Alumni.ID, &Alumni.NIM, &Alumni.Nama, &Alumni.Jurusan, &Alumni.Angkatan, &Alumni.Tahun_lulus, &Alumni.Email, &Alumni.No_telp, &Alumni.CreatedAt, &Alumni.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, Alumni)
	}
	return result, nil
}

func (r *alumniRepo) FindByID(id int) ([]models.Alumni, error) {
	rows, err := r.db.Query("SELECT id, nama, email FROM alumni WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alm []models.Alumni
	for rows.Next() {
		var alumni models.Alumni
		if err := rows.Scan(&alumni.ID, &alumni.Nama, &alumni.Email); err != nil {
			return nil, err
		}
		alm = append(alm, alumni)
	}
	return alm, nil
}


func (r *alumniRepo) Update( alumni *models.Alumni) error {
	query := `
		UPDATE alumni SET nama = $1, updated_at=$2 WHERE id=$3`
	_, err := r.db.Exec(query,
			alumni.Nama, time.Now(), alumni.ID,
	)
	return err
}


func (r *alumniRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM alumni WHERE id=$1", id)
	return err
}
