package repository

import (
	"database/sql"

	"github.com/grazierShahid/microserive-blogspot/user-service/models"
)

type User interface {
	Register(user *models.User) error
	GetByUsername(username string) (*models.User, error)
}

type user struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) User {
	return &user{db: db}
}

func (u *user) Register(user *models.User) error {
	query := `INSERT INTO users (name, email, username, password) VALUES($1, $2, $3, $4) RETURNING id, created_at`
	err := u.db.QueryRow(query, user.Name, user.Email, user.Username, user.Password).Scan(&user.ID, &user.CreatedAt)
	return err
}

func (u *user) GetByUsername(username string) (*models.User, error) {
	query := `SELECT id, name, email, username, password, active, is_deleted, is_verified, created_at, updated_at FROM users WHERE username=$1`
	row := u.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(
		&user.ID, &user.Name, &user.Email, &user.Username,
		&user.Password, &user.Active, &user.IsDeleted,
		&user.IsVerified, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
