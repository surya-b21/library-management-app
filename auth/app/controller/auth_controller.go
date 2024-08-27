package controller

import (
	"auth/app/helper"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AuthController struct
type AuthController struct{}

// SignInBody for sign in body
type SignInBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUp function
func (handler *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {

}

// SignIn function
func (handler *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	var body SignInBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// validate user later in db

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "test",
		"exp":  time.Now().Add(time.Hour * 6).Unix(),
	})

	signedToken, err := token.SignedString([]byte(helper.SigningKey))
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(signedToken))
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(helper.Salt)))
}
