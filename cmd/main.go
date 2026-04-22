package main

import (
	"backend/internal/postgres"
	"backend/internal/repository/storage/login"
	"backend/internal/repository/storage/test"
	"backend/internal/service/loginservice"
	"backend/internal/service/serviceT"
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
	//TODO принимать в user.New по interface
	testStorage := test.New(pg.DB)
	loginStorage := login.New(pg.DB)

	//service
	testService := serviceT.New(testStorage)
	loginService := loginservice.NewService(loginStorage)

	handler1 := http_transport.NewTestHandler(testService)
	handler2 := http_transport.NewLoginHandler(loginService)

	//server
	server := http_transport.New(cfg.Server)

	if err = server.Run(); err != nil {
		log.Fatal(fmt.Errorf("server run error: %w", err))
	}
}
