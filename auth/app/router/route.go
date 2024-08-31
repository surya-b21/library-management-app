package router

import (
	"net/http"

	"github.com/surya-b21/library-management-app/auth/app/controller/auth"
	"github.com/surya-b21/library-management-app/auth/app/controller/author"
	"github.com/surya-b21/library-management-app/auth/app/controller/book"
	"github.com/surya-b21/library-management-app/auth/app/controller/category"
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
	mux.HandleFunc("GET /book/{id}", middleware.UserIdentify(book.BookIdGet))
	mux.HandleFunc("PUT /book/{id}", middleware.UserIdentify(book.BookPut))
	mux.HandleFunc("DELETE /book/{id}", middleware.UserIdentify(book.BookDelete))
	mux.HandleFunc("POST /book/borrow/{id}", middleware.UserIdentify(book.BookBorrow))
	mux.HandleFunc("POST /book/return/{id}", middleware.UserIdentify(book.BookReturn))

	// author api
	mux.HandleFunc("GET /author", middleware.UserIdentify(author.AuthorGet))
	mux.HandleFunc("POST /author", middleware.UserIdentify(author.AuthorPost))
	mux.HandleFunc("GET /author/{id}", middleware.UserIdentify(author.AuthorIdGet))
	mux.HandleFunc("PUT /author/{id}", middleware.UserIdentify(author.AuthorPut))
	mux.HandleFunc("DELETE /author/{id}", middleware.UserIdentify(author.AuthorDelete))

	// category api
	mux.HandleFunc("GET /category", middleware.UserIdentify(category.CategoryGet))
	mux.HandleFunc("POST /category", middleware.UserIdentify(category.CategoryPost))
	mux.HandleFunc("GET /category/{id}", middleware.UserIdentify(category.CategoryIdGet))
	mux.HandleFunc("PUT /category/{id}", middleware.UserIdentify(category.CategoryPut))
	mux.HandleFunc("DELETE /category/{id}", middleware.UserIdentify(category.CategoryDelete))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test route"))
	})

	return mux
}
