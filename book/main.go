package main

import (
	"book/app/handler"
	"book/app/pb"
	"book/app/repo"
	"book/app/service"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	service.InitDB()

	listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	fmt.Println("Server listening on port :" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	bookRepo := &repo.BookRepository{}
	handler := handler.NewHandler(bookRepo)

	pb.RegisterBookServiceServer(s, handler)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
