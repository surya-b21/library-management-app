package main

import (
	"auth/app/router"
	"auth/app/service"
	"log"
)

func main() {
	routes := router.NewRoute()
	server := service.Server{}

	service.InitDB()

	if err := server.Run("8080", routes.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
