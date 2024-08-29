package router

import (
	"net/http"

	"github.com/surya-b21/library-management-app/auth/app/controller/auth"
	"github.com/surya-b21/library-management-app/auth/app/controller/book"
	"github.com/surya-b21/library-management-app/auth/app/middleware"
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

	// book api
	mux.HandleFunc("GET /book", middleware.UserIdentify(book.BookGet))
	mux.HandleFunc("POST /book", middleware.UserIdentify(book.BookPost))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test route"))
	})

	return mux
}
