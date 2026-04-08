package http_transport

import (
	"backend/internal/config"
	"backend/internal/transport/http_transport/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin          *gin.Engine
	cfg          config.ServerConfig
	testHandler  handler.TestHandler
	loginHandler handler.LoginHandler
}

func New(cfg config.ServerConfig,
	testHandler handler.TestHandler,
	loginHandler handler.LoginHandler) *Server {
	s := &Server{
		gin:          gin.Default(),
		cfg:          cfg,
		testHandler:  testHandler,
		loginHandler: loginHandler,
	}
	s.registerRoutes()

	return s
}
func (s *Server) registerRoutes() {
	test := s.gin.Group("/api/tests")
	{
		test.GET("/:id", s.testHandler.GetTestByID)
		test.GET("/user/:id_user", s.testHandler.GetAvailableTests)
	}
	login := s.gin.Group("/login")
	{
		login.POST("/signIn", s.loginHandler.SignIn)
	}
}
func (s *Server) Run() error {
	return s.gin.Run(":" + s.cfg.Address)
}
