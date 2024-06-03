package main

import (
	"log"

	"github.com/duvrdx/oasys-server/config"
	ginRouter "github.com/duvrdx/oasys-server/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := ginRouter.SetupRoutes()

	config.ConnectDatabase()
	router.Run()
}
