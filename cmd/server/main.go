package main

import (
	"log"

	"github.com/Msik/h-microservice/internal/pkg/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal("failed to initialize app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal("failed to run app: %v", err)
	}
}
