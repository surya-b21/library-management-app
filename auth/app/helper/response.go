package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewErrorResponse for error response
func NewErrorResponse(w http.ResponseWriter, statusCode int, errors string) {
	response, _ := json.Marshal(map[string]interface{}{
		"message": errors,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(statusCode)
	w.Write(response)
	fmt.Printf("error: %s", errors)
}

// NewSuccessResponse for success response
func NewSuccessResponse(w http.ResponseWriter, response []byte) {
	w.WriteHeader(200)
	w.Write(response)
}
