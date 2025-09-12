package repository

import (
	"crud-app/app/models"
	"database/sql"
)

type PekerjaanAlumniRepository interface {
	Create(job *models.Pekerjaan_Alm) error
	FindByAlumniID(alumniID int) ([]models.Pekerjaan_Alm, error)
	Update(job *models.Pekerjaan_Alm) error
	Delete(id int) error
}

type pekerjaanAlumniRepo struct {
	db *sql.DB
}

func NewPekerjaanAlumniRepository(db *sql.DB) PekerjaanAlumniRepository {
	return &pekerjaanAlumniRepo{db}
}


func (r *pekerjaanAlumniRepo) Create(job *models.Pekerjaan_Alm) error {
	query := `
		INSERT INTO pekerjaan_alumni
		(alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`
	return r.db.QueryRow(
		query,
		job.Alumni_ID, job.Nama_Perusahaan, job.Posisi_jabatan, job.Bidang_industri,
		job.Lokasi_kerja, job.Gaji_range, job.Tanggal_kerja, job.Tanggal_selesai,
		job.Status, job.CreatedAt, job.UpdatedAt,
	).Scan(&job.ID)
}

func (r *pekerjaanAlumniRepo) FindByAlumniID(alumniID int) ([]models.Pekerjaan_Alm, error) {
	rows, err := r.db.Query("SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, created_at, updated_at FROM pekerjaan_alumni WHERE alumni_id=$1", alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Pekerjaan_Alm
	for rows.Next() {
		var job models.Pekerjaan_Alm
		if err := rows.Scan(
			&job.ID, &job.Alumni_ID, &job.Nama_Perusahaan, &job.Posisi_jabatan, &job.Bidang_industri,
			&job.Lokasi_kerja, &job.Gaji_range, &job.Tanggal_kerja, &job.Tanggal_selesai,
			&job.Status, &job.CreatedAt, &job.UpdatedAt,
		); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *pekerjaanAlumniRepo) Update(job *models.Pekerjaan_Alm) error {
	query := `
		UPDATE pekerjaan_alumni SET nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3,
		lokasi_kerja=$4, gaji_range=$5, tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, 
		status_pekerjaan=$8, updated_at=$9 WHERE id=$10`
	_, err := r.db.Exec(query,
		job.Nama_Perusahaan, job.Posisi_jabatan, job.Bidang_industri,
		job.Lokasi_kerja, job.Gaji_range, job.Tanggal_kerja, job.Tanggal_selesai,
		job.Status, job.UpdatedAt, job.ID,
	)
	return err
}

func (r *pekerjaanAlumniRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM pekerjaan_alumni WHERE id=$1", id)
	return err
}

