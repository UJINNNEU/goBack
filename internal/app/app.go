package app

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"

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

func New(cfg *config.Config) (*App, error) {

	dbConfig := cfg.DB

	database, err := db.NewPostgresConnection(dbConfig)

	if err != nil {
		return nil, fmt.Errorf("DB err %w", err)
	}

	/*/userRepo := repository.NewUserRepository(database)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	*/

	/*loginRepo := repository.NewLoginRepository(database)

	loginService := service.NewLoginService(loginRepo)
	loginHandler := handler.NewHandler(loginService)*/


	testRepo := repository.NewTestRepository(database)
	testService := service.NewTestService(testRepo)
	testHandler := handler.NewTestHandler(testService)

	router := gin.Default()

	testHandler.RegisterRoutes(router)
	return &App{
		Router: router,
		DB:     database,
	}, nil
}
