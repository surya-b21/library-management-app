package service

import "net/http"

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

	return s.server.ListenAndServe()
}
