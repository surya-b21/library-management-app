package helper

import (
	"log"
	"net/http"
)

// NewErrorResponse for error response
func NewErrorResponse(w http.ResponseWriter, statusCode int, errors string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(errors))
	log.Printf("error: %s", errors)
}

// NewSuccessResponse for success response
func NewSuccessResponse(w http.ResponseWriter, messages string) {
	w.WriteHeader(200)
	w.Write([]byte(messages))
}
