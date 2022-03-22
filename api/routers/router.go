package routers

import (
	"github.com/gin-gonic/gin"
	api_jwt "github.com/smartschool/api/api-jwt"
	api_device "github.com/smartschool/api/routers/api-device"
	api_mobile "github.com/smartschool/api/routers/api-mobile"
)

func Initialize() (*gin.Engine, error) {
	r := gin.New()

	authMw, err := api_jwt.GetDefaultGinJWTMiddleware()
	if err != nil {
		return nil, err
	}

	mobileGroup := r.Group("/mobile")
	mobileGroup.Use(authMw.GetAuthFunc())

	mobileGroup.GET("/late/:id", api_device.GetStudentCheckInLateHistory)
	mobileGroup.PUT("/change-password", api_mobile.UpdatePassword)

	r.POST("/checkin", api_device.EventCheckin)
	r.POST("/login", api_mobile.Login)
	//r.GET("/courses/:id", api_device.GetCourseByID)
	//r.GET("/courses", api_device.GetCourses)

	return r, nil
}
