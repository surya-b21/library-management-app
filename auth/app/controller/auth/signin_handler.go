package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/model"
	"github.com/surya-b21/library-management-app/auth/app/service"

	"github.com/golang-jwt/jwt/v5"
)

// SignIn function
func (handler *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var body SignInBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// validate user
	db := service.DB
	var user model.User
	mod := db.Where("username = ?", body.Username).First(&user)

	if mod.RowsAffected < 1 {
		helper.NewErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	hashPassword := generatePasswordHash(body.Password)

	if hashPassword != *user.Password {
		helper.NewErrorResponse(w, http.StatusBadRequest, "Wrong user / password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "test",
		"exp":  time.Now().Add(time.Hour * 6).Unix(),
	})

	signedToken, err := token.SignedString([]byte(helper.SigningKey))
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	response, _ := json.Marshal(map[string]interface{}{
		"token":   signedToken,
		"expired": 6 * 3600,
	})

	helper.NewSuccessResponse(w, response)
}
