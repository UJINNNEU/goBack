package handler

import (
	_"log"
	"net/http"
	_"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerHello(c *gin.Context) {
	c.String(http.StatusOK, "Server opa!")
}

func HandlerGetTest(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(500,gin.H{
		"message-error": "не смогли преобразовать id в число, наш косяк"})
	}
	c.JSON(http.StatusOK, gin.H{
	"message":"get test by id ",
	"id:":id})
}