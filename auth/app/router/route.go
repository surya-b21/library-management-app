package router

import (
	"auth/app/controller/auth"
	"auth/app/middleware"
	"net/http"
)

// Route struct
type Route struct{}

// NewRoute function
func NewRoute() *Route {
	return &Route{}
}

// InitRoutes to initiate route
func (r *Route) InitRoutes() *http.ServeMux {
	auth := auth.AuthHandler{}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sign-up", auth.SignUp)
	mux.HandleFunc("POST /sign-in", auth.SignIn)
	mux.HandleFunc("GET /protected", middleware.UserIdentify(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success access protected route"))
	}))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test route"))
	})

	return mux
}
