package api_mobile

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	api_jwt "github.com/smartschool/api/api-jwt"
	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/service"
	"golang.org/x/crypto/bcrypt"
)

// func Register(c *gin.Context) {
// 	var registerReq dto.RegisterRequest
// 	err := c.ShouldBindJSON(&registerReq)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, errors.New("register request is invalid"))
// 		return
// 	}
// 	isMail, email := helper.CheckMailFormat(registerReq.Email)
// 	if !isMail {
// 		c.JSON(http.StatusBadRequest, errors.New("wrong email request"))
// 		return
// 	}

// 	//password format error
// 	if len(registerReq.Password) < 8 {
// 		c.JSON(http.StatusBadRequest, errors.New("password must have at least 8 characters"))
// 	}

// 	passwordHash, err := helper.HashPassword(registerReq.Password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, errors.New("password hash fail"))
// 		return
// 	}

// 	user := entity.User{
// 		Email:    email,
// 		Password: passwordHash,
// 	}
// 	database.DbInstance.Create(&user)
// 	c.JSON(http.StatusOK, "Register success")
// }

func Login(c *gin.Context) {
	var request dto.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Login request is invalid",
		})
		return
	}

	var user entity.User
	err = database.DbInstance.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong password",
		})
		return
	}

	authMw, err := api_jwt.GetDefaultGinJWTMiddleware()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
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

func UpdatePassword(c *gin.Context) {
	var req = dto.UpdatePasswordRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messgae": "Update Password request is invalid",
		})
		return
	}
	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get userID"})
		return
	}
	err = service.UpdatePassword(fmt.Sprint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "2 password not matches"})
		return
	}

	c.JSON(http.StatusOK, req)
}
