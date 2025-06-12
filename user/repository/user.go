package repository

import "database/sql"

type user struct {
	db *sql.DB
}

type User interface{
	
}

func NewUserRepo(db *sql.DB) User {
	return &user{db: db}
}
