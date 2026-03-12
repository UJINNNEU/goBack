package main

import (
	"backend/internal/handler"
	_ "fmt"
	_ "log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.HandlerHello)
	router.GET("/test/:id", handler.HandlerGetTest)

	router.Run()

}
