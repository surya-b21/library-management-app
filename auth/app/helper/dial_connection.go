package helper

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = map[string]string{
	"book": "book-service:8080",
}

func ServerDial(service string) *grpc.ClientConn {
	serverAddr := flag.String(
		"server", address[service],
		"The server address in the format of host:port",
	)
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return nil
	}

	return conn
}
