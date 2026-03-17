package main

import (
	_ "backend/internal/handler"
	"fmt"
	_ "fmt"
	_ "log"

	"backend/internal/app"
)

func main() {
	app, err := app.New()

	if err != nil {
		fmt.Println("Start!!!")
	}

	if err := app.Run(":8080"); err != nil {
		fmt.Println("Start!!!")
	}
}
