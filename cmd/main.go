package main

import (
	"backend/internal/postgres"
	"backend/internal/repository/storage/user"
	userSrv "backend/internal/service/user"
	"backend/internal/transport/http_transport"
	"fmt"
	"log"

	"backend/internal/config"
)

func main() {
	//configs
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	app, err := app.New(cfg)

	//entities
	pg, err := postgres.New(cfg.DB)
	if err != nil {
		fmt.Println("Start!!!")
	}

	if err := app.Run(":8080"); err != nil {
		fmt.Println("Start!!!")
	}

}
