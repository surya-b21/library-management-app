package main

import (
	"auth/app/router"
	"auth/app/service"
	"log"
	"os"
)

func main() {
	routes := router.NewRoute()
	server := service.Server{}

	service.InitDB()

	if err := server.Run(os.Getenv("PORT"), routes.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
