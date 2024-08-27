package auth

import (
	"auth/app/helper"
	"auth/app/model"
	"auth/app/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SignUp function
func (handler *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var body SignUpBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	hashPassword := generatePasswordHash(body.Password)
	db := service.DB

	newUser := model.User{}
	newUser.Name = &body.Name
	newUser.Username = &body.Username
	newUser.Password = &hashPassword
	db.Create(&newUser)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": newUser.ID,
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