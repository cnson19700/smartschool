package api_mobile

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	api_jwt "github.com/smartschool/api/api-jwt"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/model/dto"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var request dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.New("login request is invalid"))
	}

	var user entity.User
	err = database.DbInstance.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.New("user not found"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.New("wrong password"))
	}

	authMw, err := api_jwt.GetDefaultGinJWTMiddleware()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	// Create JWT token
	token := jwt.New(jwt.GetSigningMethod("RS512"))
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	expire := time.Now().Add(time.Hour * 24 * 30 * 12)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()

	tokenString, _ := authMw.GetSignedString(token)
	resp := map[string]interface{}{"token": tokenString}

	c.JSON(http.StatusOK, resp)
}
