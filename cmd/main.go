package main

import (
	"backend/internal/postgres"
	"backend/internal/repository/storage/test"
	"backend/internal/repository/storage/login"
	"backend/internal/service/serviceT"
	"backend/internal/service/loginservice"
	"backend/internal/transport/http_transport"
	"backend/internal/transport/http_transport/handler"
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
	loginStorage :=	login.New(pg.DB)

	//service
	testService := serviceT.New(testStorage)
	loginService := loginservice.NewService(loginStorage)

	handler1 := handler.NewTestHandler(testService)
	handler2 := handler.NewLoginHandler(loginService)

	//server
	server := http_transport.New(cfg.Server, *handler1, *handler2)

	if err = server.Run(); err != nil {
		log.Fatal(fmt.Errorf("server run error: %w", err))
	}
}
