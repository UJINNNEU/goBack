package app

import (
	"backend/internal/db"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	_"backend/internal/config"

	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func (a *App) Run(addres string) error {
	return a.Router.Run(addres)
}

func New() (*App, error) {

	dbConfig := db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "123456",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	database, err := db.NewPostgresConnection(dbConfig)

	if err != nil {
		return nil, fmt.Errorf("BaseData govno %w", err)
	}

	// 2. Инициализация репозитория
	userRepo := repository.NewUserRepository(database)

	// 3. Инициализация сервиса
	userService := service.NewUserService(userRepo)

	// 4. Инициализация хендлера
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	userHandler.RegisterRoutes(router)
	return &App{
		Router: router,
		DB:     database,
	}, nil
}
