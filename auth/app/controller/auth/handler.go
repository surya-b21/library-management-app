package auth

import (
	"crypto/sha256"
	"fmt"

	"github.com/surya-b21/library-management-app/auth/app/helper"
)

// AuthHandler struct
type AuthHandler struct{}

// SignInBody for sign in body
type SignUpBody struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignInBody for sign in body
type SignInBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(helper.Salt)))
}
