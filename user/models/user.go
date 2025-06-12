package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Active     bool      `json:"active"`
	IsDeleted  bool      `json:"is_deleted"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
