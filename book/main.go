package main

import (
	"book/app/handler"
	"book/app/pb"
	"book/app/repo"
	"book/app/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	service.InitDB()

	listener, err := net.Listen("tcp", ":8080")
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
