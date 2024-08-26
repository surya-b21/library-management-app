package main

import (
	"auth/app/router"
	"auth/app/service"
	"log"
)

func main() {
	routes := router.NewRoute()
	server := new(service.Server)

	if err := server.Run("8000", routes.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
