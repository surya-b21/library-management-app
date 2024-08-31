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

// SignUp function
// @Summary      Sign Up
// @Description  Sign up new account
// @Tags         Auth
// @Accept       application/json
// @Produce		 application/json
// @Param        data   body  auth.SignUpBody  true  "Sign Up Payload"
// @Success      200  {object} auth.JsonResult
// @Router       /sign-up [post]
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
