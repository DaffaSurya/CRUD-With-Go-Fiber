package handlers

import (
	"crud-app/database"
	"crud-app/app/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// CREATE (POST) table alumni

func CreateUserAlumni(c *fiber.Ctx) error {
	var Alumni models.Alumni
	if err := c.BodyParser(&Alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	query := "INSERT INTO alumni (nim, nama , jurusan, angkatan,  tahun_lulus, email, no_telepon, alamat , created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	err := database.DB.QueryRow(query, Alumni.NIM, Alumni.Nama, Alumni.Jurusan, Alumni.Angkatan, Alumni.Tahun_lulus, Alumni.Email, Alumni.No_telp, Alumni.Alamat, time.Now(),time.Now()).Scan(&Alumni.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(Alumni)
}

// Get User table alumni
func GetUsersAlumni(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, created_dated_atat, up  FROM alumni ORDER BY created_at DESC")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	users := []models.Alumni{}
	for rows.Next() {
		var Alumni models.Alumni
		if err := rows.Scan(&Alumni.ID, &Alumni.NIM, &Alumni.Nama, &Alumni.Jurusan, &Alumni.Angkatan, &Alumni.Tahun_lulus, &Alumni.Email, &Alumni.No_telp, &Alumni.CreatedAt, &Alumni.UpdatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		users = append(users, Alumni)
	}

	return c.JSON(users)
}

// Get data Alumni berdasarkan ID

func GetUsersAlumniID(c *fiber.Ctx) error {
	id := c.Params("id")
	var Alumni models.Alumni
	err := database.DB.QueryRow("SELECT id, nama, email FROM alumni WHERE id = $1", id).Scan(&Alumni.ID, &Alumni.Nama, &Alumni.Email)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "data tidak ditemukan"})
	}
	return c.JSON(Alumni)
}


// Get data Alumni berdasarkan ID

// Update User Alumni
func UpdateUserAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var Alumni models.Alumni
	if err := c.BodyParser(&Alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// query := "UPDATE alumni SET nama = $1, jurusan = $2, angkatan = $3, email = $4, updated_at = $5 WHERE id = $6"
    query := "UPDATE alumni SET nama = $1,  updated_at = $2 WHERE id = $3"
	// _, err := database.DB.Exec(query, Alumni.Nama, Alumni.Jurusan, Alumni.Angkatan, Alumni.Email, time.Now(), id)
    _, err := database.DB.Exec(query, Alumni.Nama,  time.Now(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus data alumni",
		})
	}

	Alumni.ID = atoi(id)
	return c.JSON(Alumni)
}

// Update User Alumni

func DeleteUserAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM alumni WHERE id = $1", id)
	if err != nil {
		// return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus data alumni",
		})
	}
	return c.SendStatus(204)
}


// CreatePekerjaan_alumni

// func CreatePekerjaan_alumni(c *fiber.Ctx) error {
// 	var Job models.Pekerjaan_Alm
// 	if err := c.BodyParser(&Job); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	query := "INSERT INTO pekerjaan_alumni (nama_perusahaan , posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
// 	err := database.DB.QueryRow(query, Job.Nama_Perusahaan, Job.Posisi_jabatan, Job.Bidang_industri, Job.Lokasi_kerja, Job.Gaji_range, Job.Tanggal_kerja, Job.Tanggal_selesai, Job.Status, time.Now(),time.Now()).Scan(&Job.ID)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(201).JSON(Job)
// }


func CreatePekerjaanAlumni(c *fiber.Ctx) error {
	var job models.Pekerjaan_Alm
	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Query INSERT
	query := `
		INSERT INTO pekerjaan_alumni 
		(alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		job.Alumni_ID,
		job.Nama_Perusahaan,
		job.Posisi_jabatan,
		job.Bidang_industri,
		job.Lokasi_kerja,
		job.Gaji_range,
		job.Tanggal_kerja,
		job.Tanggal_selesai,
		job.Status,
		time.Now(),
		time.Now(),
	).Scan(&job.ID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(job)
}


// GetPekerjaan_alumni

// Get data dari tabel pekerjaan_alumni
func GetPekerjaanAlumni(c *fiber.Ctx) error {
    rows, err := database.DB.Query(`
        SELECT 
            id, 
            alumni_id, 
            nama_perusahaan, 
            posisi_jabatan, 
            bidang_industri, 
            lokasi_kerja, 
            gaji_range, 
            tanggal_mulai_kerja, 
            status_pekerjaan, 
            deskripsi_pekerjaan, 
            created_at, 
            updated_at
        FROM pekerjaan_alumni
    `)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    defer rows.Close()

    users := []models.Pekerjaan_Alm{}
    for rows.Next() {
        var pekerjaan models.Pekerjaan_Alm
        if err := rows.Scan(
            &pekerjaan.ID,
            &pekerjaan.Alumni_ID,
            &pekerjaan.Nama_Perusahaan,
            &pekerjaan.Posisi_jabatan,
            &pekerjaan.Bidang_industri,
            &pekerjaan.Lokasi_kerja,
            &pekerjaan.Gaji_range,
            &pekerjaan.Tanggal_kerja,
            &pekerjaan.Status,
            &pekerjaan.Deskripsi,
            &pekerjaan.CreatedAt,
            &pekerjaan.UpdatedAt,
        ); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }
        users = append(users, pekerjaan)
    }

    return c.JSON(users)
}

// Get pekerjaan alumni by ID

func GetPekerjaan_alumni_ID(c *fiber.Ctx) error {
	id := c.Params("id")
	var pekerjaan models.Pekerjaan_Alm
	err := database.DB.QueryRow("SELECT id, nama_perusahaan, posisi_jabatan FROM pekerjaan_alumni WHERE id = $1", id).Scan(&pekerjaan.ID, &pekerjaan.Nama_Perusahaan, &pekerjaan.Posisi_jabatan)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "data tidak ditemukan"})
	}
	return c.JSON(pekerjaan)
}

// Get pekerjaan alumni by ID


// Update pekerjaan alumni

func UpdatePekerjaanAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var Job models.Pekerjaan_Alm
	if err := c.BodyParser(&Job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// query := "UPDATE alumni SET nama = $1, jurusan = $2, angkatan = $3, email = $4, updated_at = $5 WHERE id = $6"
    query := "UPDATE pekerjaan_alumni SET nama_perusahaan = $1, posisi_jabatan = $2,  updated_at = $3 WHERE id = $4"
	_, err := database.DB.Exec(query, Job.Nama_Perusahaan, Job.Posisi_jabatan,  time.Now(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus data alumni",
		})
	}

	Job.ID = atoi(id)
	return c.JSON(Job)
}



// menghapus pekerjaan alumni

func DeletePekerjaan_alumni(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM pekerjaan_alumni WHERE id = $1", id)
	if err != nil {
		// return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus data Pekerjaan alumni",
		})
	}
	return c.SendStatus(204)
}


// Get pekerjaan alumni by Alumni - id

func GetPekerjaanByAlumni(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")

	rows, err := database.DB.Query(`
		SELECT id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, 
		       gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, 
		       created_at, updated_at
		FROM pekerjaan_alumni
		WHERE alumni_id = $1
	`, alumniID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var pekerjaanList []models.Pekerjaan_Alm

	for rows.Next() {
		var pekerjaan models.Pekerjaan_Alm
		if err := rows.Scan(
			&pekerjaan.ID,
			&pekerjaan.Nama_Perusahaan,
			&pekerjaan.Posisi_jabatan,
			&pekerjaan.Bidang_industri,
			&pekerjaan.Lokasi_kerja,
			&pekerjaan.Gaji_range,
			&pekerjaan.Tanggal_kerja,
			&pekerjaan.Tanggal_selesai,
			&pekerjaan.Status,
			&pekerjaan.CreatedAt,
			&pekerjaan.UpdatedAt,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		pekerjaanList = append(pekerjaanList, pekerjaan)
	}

	return c.JSON(pekerjaanList)
}















