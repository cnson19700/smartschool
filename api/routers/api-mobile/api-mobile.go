package api_mobile

import (
	"fmt"
	"net/http"
	"time"

	"github.com/smartschool/service/fireapp"

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
	err = database.DbInstance.Where("user_name = ?", request.Email).First(&user).Error
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
	claims["faculty_id"] = user.FacultyID

	tokenString, _ := authMw.GetSignedString(token)
	resp := map[string]interface{}{
		"username":    user.LastName + " " + user.FirstName,
		"is_activate": user.IsActivate,
		"token":       tokenString,
	}

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
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	err = service.UpdatePassword(fmt.Sprint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, req)
}

func GetMe(c *gin.Context) {
	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	res, err := service.GetMe(fmt.Sprint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get student profile"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetCourseAttendanceOfOneUser(c *gin.Context) {
	request := struct {
		CourseID uint `form:"course_id" binding:"required"`
	}{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot capture request",
		})
		return
	}

	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "cannot find this user"})
		return
	}
	userId, canConvert := id.(float64)
	if !canConvert {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authenticate fail"})
		return
	}

	course, err := service.GetCourseBasicInfoByID(request.CourseID)
	if err != nil || course == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot verified course",
		})
		return
	}

	res, err := service.GetAttendanceInCourseOneUser(request.CourseID, uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get course attendance for user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course":          course.CourseID + " - " + course.Name,
		"attendance_list": res,
	})
}

func GetInDayAttendance(c *gin.Context) {
	request := struct {
		TimezoneOffset int `form:"time_offset" binding:"required"`
	}{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot capture request",
		})
		return
	}

	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	userId, canConvert := id.(float64)
	if !canConvert {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authenticate fail"})
		return
	}

	faculty_id, isGet := c.Get("facultyId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	userFacultyId, canConvert := faculty_id.(float64)
	if !canConvert {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authenticate fail"})
		return
	}

	res, err := service.GetCheckInHistoryInDay(uint(userId), uint(userFacultyId), request.TimezoneOffset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get attendance history for user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"checkin_list": res,
	})
}

func GetQREncodeString(c *gin.Context) {
	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	userId, canConvert := id.(float64)
	if !canConvert {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authenticate fail"})
		return
	}

	res, err := service.GenerateQREncodeString(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot provide QR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"qr_string": res,
	})
}

func UpdateNotificationToken(c *gin.Context) {
	id, _ := c.Get("userId")
	userId, _ := id.(float64)

	var request dto.UpdateNotificationTokenRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Notification token update request is invalid",
		})
		return
	}

	var userNotificationToken entity.UserNotificationToken
	database.DbInstance.First(&userNotificationToken, "token = ?", request.NotificationToken)

	if userNotificationToken.ID == 0 {
		newUserNotificationToken := entity.UserNotificationToken{
			UserID: uint(userId),
			Token:  request.NotificationToken,
		}

		err = database.DbInstance.Create(&newUserNotificationToken).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}
	}
}

func GetCourseInSemesterOfOneUser(c *gin.Context) {
	request := struct {
		SemesterID uint `form:"semester_id" binding:"required"`
	}{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot capture request",
		})
		return
	}

	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	userId, canConvert := id.(float64)
	if !canConvert {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authenticate fail"})
		return
	}

	res, err := service.GetListCourseByUserSemester(uint(userId), request.SemesterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot get list course in selected semester for user",
		})
		return
	}

	if res == nil {
		res = make([]dto.CourseReportListElement, 0)
	}
	c.JSON(http.StatusOK, gin.H{
		"course_list": res,
	})
}

func GetSemesterInFaculty(c *gin.Context) {
	faculty_id, isGet := c.Get("facultyId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cannot get userID"})
		return
	}
	userFacultyId, canConvert := faculty_id.(float64)
	if !canConvert {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authenticate fail"})
		return
	}

	res, err := service.GetSemesterByFacultyID(uint(userFacultyId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot get list semester for user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"semester_list": res,
	})
}

func ChangePasswordFirstTime(c *gin.Context) {
	var req = dto.ChangePasswordFirstTimeRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"messgae":     "Update Password request is invalid",
			"is_activate": false,
		})
		return
	}
	id, isGet := c.Get("userId")
	if !isGet {
		c.JSON(http.StatusNotFound, gin.H{
			"message":     "Cannot get userID",
			"is_activate": false,
		})
		return
	}
	res, err := service.ChangePasswordFirstTime(fmt.Sprint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprint(err),
			"is_activate": res,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Mật khẩu đã được thay đổi",
		"is_activate": res,
	})
}

func TestNotification(c *gin.Context) {
	id, _ := c.Get("userId")
	userId, _ := id.(float64)

	data := map[string]string{
		"message": "Hello world",
	}

	err := fireapp.SendNotification(uint(userId), data)
	_ = err
	c.JSON(http.StatusOK, data)
}

func ResetPassword(c *gin.Context) {
	var request dto.ResetPasswordRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Reset password request is invalid",
		})
		return
	}

	err = service.ResetPassword(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password reset email is sent",
	})
}
