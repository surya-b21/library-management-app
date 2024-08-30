package helper

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = map[string]string{
	"book":     "book-service:8080",
	"author":   "author-service:8080",
	"category": "category-service:8080",
}

func ServerDial(service string) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(address[service], opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return nil
	}

	return conn
}
