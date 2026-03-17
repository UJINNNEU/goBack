package app

import (
	"backend/internal/db"
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
	router := gin.Default()

	return &App{
		Router: router,
		DB:     database,
	}, nil
}
