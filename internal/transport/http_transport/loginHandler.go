package http_transport

import (
	"backend/internal/model"
	"backend/internal/service/loginservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Login    string `json:"login"`
	Password int    `json:"password"`
}

type LoginHandler struct {
	loginService *loginservice.LoginService
}

func NewLoginHandler(loginService *loginservice.LoginService) *LoginHandler {
	return &LoginHandler{
		loginService: loginService}
}

func (l *LoginHandler) SignIn(c *gin.Context) {

	var requestBody requestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"login": "fail",
			"error": err.Error()})
		return
	}
	var request model.LoginRequest = model.LoginRequest{
		Login:    requestBody.Login,
		Password: requestBody.Password,
	}
	response, err := l.loginService.SignIn(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"login": "fail",
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"login":    "success",
		"responce": response})
	return
}
