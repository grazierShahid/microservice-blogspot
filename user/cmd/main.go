package main

import (
	"github.com/grazierShahid/microserive-blogspot/user-service/config"
	"github.com/grazierShahid/microserive-blogspot/user-service/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)
	defer db.Close()

	repo := repository.NewUserRepo(db)
	svc := service.NewUserService(repo)
}
