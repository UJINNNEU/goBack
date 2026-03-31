package main

import (
	"fmt"
	_ "log"

	"backend/internal/app"
	"backend/internal/config"
)

func main() {

	fmt.Println("Start")
	cfg, err := config.Load()

	if err != nil {
		fmt.Print(err)
		return
	}
	app, err := app.New(cfg)

	defer app.DB.Close()

	if err != nil {
		fmt.Println(err)
	}

	if err := app.Run(":8084"); err != nil {
		fmt.Println(err)
	}

}
