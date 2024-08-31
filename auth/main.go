package main

import (
	"log"
	"os"

	"github.com/surya-b21/library-management-app/auth/app/router"
	"github.com/surya-b21/library-management-app/auth/app/service"
)

// @title Library Management App API
// @version 1.0
// @description For library management api for test purpose
// @host localhost:8080
// @BasePath /
// @securitydefinitions.apikey
// @in header
// @name Authorization
func main() {
	routes := router.NewRoute()
	server := service.Server{}

	service.InitDB()

	if err := server.Run(os.Getenv("PORT"), routes.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
