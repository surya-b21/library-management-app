package service

import (
	"fmt"
	"net/http"
)

// Server struct
type Server struct {
	server *http.Server
}

// Run server
func (s *Server) Run(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	fmt.Println("Server listening on port :" + port)
	return s.server.ListenAndServe()
}
