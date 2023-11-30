package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/app"
	"log"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	application := app.NewApp(cfg)
	application.Run()

}
