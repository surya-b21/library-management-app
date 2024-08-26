package router

import "net/http"

// Route struct
type Route struct{}

// NewRoute function
func NewRoute() *Route {
	return &Route{}
}

// InitRoutes to initiate route
func (r *Route) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test Route"))
	})

	return mux
}
