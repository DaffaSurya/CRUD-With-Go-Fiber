package repository

import (
	"crud-app/app/models"
	"database/sql"
)

type UserRepository interface {
	GetByUsername(username string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
}

type userPostgres struct{ db *sql.DB }

func NewUserRepository(db *sql.DB) UserRepository {
	return &userPostgres{db}
}

func (r *userPostgres) GetByUsername(username string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT id, username, password_hash, role FROM users WHERE username=$1", username).
		Scan(&u.ID, &u.Username, &u.Password, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userPostgres) GetByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT id, username, role FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Username, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}


func (r *userPostgres) Create(u *models.User) error {
    query := `INSERT INTO users (email, username, password_hash, role) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(query, u.Email, u.Username, u.Password, u.Role)
    return err
}

