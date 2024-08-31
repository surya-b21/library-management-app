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

	unauthorizedRoute := http.NewServeMux()
	unauthorizedRoute.HandleFunc("POST /sign-up", auth.SignUp)
	unauthorizedRoute.HandleFunc("POST /sign-in", auth.SignIn)
	unauthorizedRoute.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test route"))
	})

	authorizedRoute := http.NewServeMux()

	// book api
	authorizedRoute.HandleFunc("GET /book", book.BookGet)
	authorizedRoute.HandleFunc("POST /book", book.BookPost)
	authorizedRoute.HandleFunc("GET /book-recomendation/", book.BookRecomendation)
	authorizedRoute.HandleFunc("POST /book/borrow/{id}", book.BookBorrow)
	authorizedRoute.HandleFunc("POST /book/return/{id}", book.BookReturn)
	authorizedRoute.HandleFunc("GET /book/{id}", book.BookIdGet)
	authorizedRoute.HandleFunc("PUT /book/{id}", book.BookPut)
	authorizedRoute.HandleFunc("DELETE /book/{id}", book.BookDelete)

	// author api
	authorizedRoute.HandleFunc("GET /author", author.AuthorGet)
	authorizedRoute.HandleFunc("POST /author", author.AuthorPost)
	authorizedRoute.HandleFunc("GET /author/{id}", author.AuthorIdGet)
	authorizedRoute.HandleFunc("PUT /author/{id}", author.AuthorPut)
	authorizedRoute.HandleFunc("DELETE /author/{id}", author.AuthorDelete)

	// category api
	authorizedRoute.HandleFunc("GET /category", category.CategoryGet)
	authorizedRoute.HandleFunc("POST /category", category.CategoryPost)
	authorizedRoute.HandleFunc("GET /category/{id}", category.CategoryIdGet)
	authorizedRoute.HandleFunc("PUT /category/{id}", category.CategoryPut)
	authorizedRoute.HandleFunc("DELETE /category/{id}", category.CategoryDelete)

	authorizedRoute.HandleFunc("GET /protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success access protected route"))
	})

	authorizedMiddleware := middleware.MiddlewareStack(
		middleware.Cors,
		middleware.UserIdentify,
	)

	router := http.NewServeMux()
	router.Handle("/", authorizedMiddleware(authorizedRoute))
	router.Handle("/auth/", http.StripPrefix("/auth", middleware.Cors(unauthorizedRoute)))
	return router
}
