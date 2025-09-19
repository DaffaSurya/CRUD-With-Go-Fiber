package repository

import (
	
	"crud-app/app/models"
	"database/sql"
)
type PekerjaanRepository interface {
	FindByAlumni(alumniID int) ([]models.Pekerjaan, error)
	FindAllPekerjaan() ([]models.Pekerjaan, error)
	Create(p *models.Pekerjaan) error
	Update(id int, p *models.Pekerjaan) error
	Delete(id int) error
}

type pekerjaanPostgres struct{ db *sql.DB }

func NewPekerjaanRepository(db *sql.DB) PekerjaanRepository {
	return &pekerjaanPostgres{db}
}

func (r *pekerjaanPostgres) FindByAlumni(alumniID int) ([]models.Pekerjaan, error) {
	rows, err := r.db.Query("SELECT id, alumni_id, nama, posisi FROM pekerjaan_alumni WHERE alumni_id=$1", alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Pekerjaan
	for rows.Next() {
		var p models.Pekerjaan
		rows.Scan(&p.ID, &p.Alumni_ID, &p.Nama_Perusahaan, &p.Posisi_jabatan)
		list = append(list, p)
	}
	return list, nil
}

func (r *pekerjaanPostgres) Create(p *models.Pekerjaan) error {
	_, err := r.db.Exec("INSERT INTO pekerjaan_alumni(alumni_id,nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, created_at, updated_at) VALUES($1,$2,$3, $4, $5, $6, $7, $8, $9, $10, $11)",
		p.Alumni_ID, p.Nama_Perusahaan, p.Posisi_jabatan, p.Bidang_industri, p.Lokasi_kerja, p.Gaji_range, p.Tanggal_kerja, p.Tanggal_selesai, p.Status, p.CreatedAt, p.UpdatedAt)
	return err
}

func (r *pekerjaanPostgres) Update(id int, p *models.Pekerjaan) error {
	_, err := r.db.Exec("UPDATE pekerjaan_alumni SET nama_perusahaan=$1, posisi_jabatan=$2 , gaji_range=$3 WHERE id=$4", p.Nama_Perusahaan, p.Posisi_jabatan, p.Gaji_range, id)
	return err
}

func (r *pekerjaanPostgres) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM pekerjaan_alumni WHERE id=$1", id)
	return err
}


func (r *pekerjaanPostgres) FindAllPekerjaan() ([]models.Pekerjaan, error) {
	rows, err := r.db.Query("SELECT id, nama_perusahaan, posisi_jabatan FROM pekerjaan_alumni")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Pekerjaan
	for rows.Next() {
		var a models.Pekerjaan
		rows.Scan(&a.ID, &a.Nama_Perusahaan, &a.Posisi_jabatan)
		list = append(list, a)
	}
	return list, nil
}