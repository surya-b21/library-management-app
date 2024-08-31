package helper

import (
	"crypto/tls"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address = map[string]string{
	"book":     "book-service:8080",
	"author":   "author-service:8080",
	"category": "category-service:8080",
}

func ServerDial(service string) *grpc.ClientConn {
	opts := []grpc.DialOption{
		// uncomment it when using in localhost
		// grpc.WithTransportCredentials(insecure.NewCredentials()),

		// use when build dockerfile
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})),
	}

	conn, err := grpc.NewClient(address[service], opts...)
	if err != nil {
		log.Printf("fail to dial: %v\n", err)
		return nil
	}

	return conn
}
