package http_transport

import (
	"backend/internal/config"
	"backend/internal/model"
	"context"

	"github.com/gin-gonic/gin"
)

type userService interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type Server struct {
	gin  *gin.Engine
	cfg  config.ServerConfig
	user userService
}

func New(cfg config.ServerConfig, user userService) *Server {
	s := &Server{
		gin:  gin.Default(),
		cfg:  cfg,
		user: user,
	}
	s.registerRoutes()

	return s
}

func (s *Server) registerRoutes() {
	users := s.gin.Group("/api/users")
	{
		users.GET("/:id", s.GetUser)
		users.GET("/", s.ListUsers)
	}
}

func (s *Server) Run() error {
	return s.gin.Run(s.cfg.Address)
}
