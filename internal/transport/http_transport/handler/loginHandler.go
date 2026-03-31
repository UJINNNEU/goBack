package handler

import (
	"backend/internal/service/login"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/model"
)

type LoginHandler struct {
	loginService service.LoginService
}

func NewHandler(loginService service.LoginService) *LoginHandler {
	return &LoginHandler{
		loginService: loginService,
	}
}

func (h *LoginHandler) RegisterRoutes(router *gin.Engine) {
	login := router.Group("/api/signIn")
	{
		login.POST("", h.SignIn)
	}
}

func (h *LoginHandler) SignIn(c *gin.Context) {

	userLogin := model.LoginRequest{}
	err := c.ShouldBindBodyWithJSON(&userLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	resp, err := h.loginService.LogIn(c.Request.Context(), userLogin)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message := fmt.Sprintf("Name %s, role %s, id %d",
		resp.Name, resp.Role, resp.Id)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	return
}
