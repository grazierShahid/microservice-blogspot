package repository

import (
	"database/sql"

	"github.com/grazierShahid/microserive-blogspot/user-service/models"
)

type User interface {
	Register(user *models.User) error
	//Login()
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
