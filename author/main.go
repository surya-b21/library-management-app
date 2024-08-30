package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/surya-b21/library-management-app/author/app/handler"
	"github.com/surya-b21/library-management-app/author/app/pb"
	"github.com/surya-b21/library-management-app/author/app/repo"
	"github.com/surya-b21/library-management-app/author/app/service"
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

	authorRepo := &repo.AuthorRepository{}
	handler := handler.NewHandler(authorRepo)

	pb.RegisterAuthorServiceServer(s, handler)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
