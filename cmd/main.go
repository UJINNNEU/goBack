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

	//entities
	pg, err := postgres.New(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	//repository
	userStorage := user.New(pg.DB)

	//service
	userService := userSrv.New(userStorage)

	//server
	server := http_transport.New(cfg.Server, userService)

	if err = server.Run(); err != nil {
		log.Fatal(fmt.Errorf("server run error: %w", err))
	}
}
