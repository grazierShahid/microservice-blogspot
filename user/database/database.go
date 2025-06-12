package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/grazierShahid/microserive-blogspot/user-service/config"
	_ "github.com/lib/pq"
)

func InitDB(cfg config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open DB: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}

	return db
}
