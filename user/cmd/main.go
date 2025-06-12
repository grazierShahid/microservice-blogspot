package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/grazierShahid/microserive-blogspot/user-service/config"
	"github.com/grazierShahid/microserive-blogspot/user-service/database"
	"github.com/grazierShahid/microserive-blogspot/user-service/handler"
	"github.com/grazierShahid/microserive-blogspot/user-service/repository"
	"github.com/grazierShahid/microserive-blogspot/user-service/service"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)
	defer db.Close()

	migrationSQL, err := os.ReadFile("/home/grazier/Desktop/microservice-project/user/migrations/0001/up.sql")
	if err != nil {
		log.Fatal("Failed to read migration file: ", err)
	}

	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatal("db migration does not executing: ", err)
	}

	log.Println("Migration executed sucessfully!")

	repo := repository.NewUserRepo(db)
	svc := service.NewUserService(repo)

	r := mux.NewRouter()
	handler.NewUserHandler(r, svc)

	fmt.Println("working!!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
