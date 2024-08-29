package main

import (
	"log"
	"os"

	"github.com/surya-b21/library-management-app/auth/app/router"
	"github.com/surya-b21/library-management-app/auth/app/service"
)

func main() {
	routes := router.NewRoute()
	server := service.Server{}

	service.InitDB()

	if err := server.Run(os.Getenv("PORT"), routes.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
