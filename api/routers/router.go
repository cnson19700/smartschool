package routers

import (
	"github.com/gin-gonic/gin"
	api_device "github.com/smartschool/api/routers/api-device"
	api_mobile "github.com/smartschool/api/routers/api-mobile"
)

func Initialize() (*gin.Engine, error) {
	r := gin.New()

	// authMw, err := api_jwt.GetDefaultGinJWTMiddleware()
	// if err != nil {
	// 	return nil, err
	// }

	mobileGroup := r.Group("/mobile")
	//mobileGroup.Use(authMw.GetAuthFunc())

	mobileGroup.GET("/late/:id", api_device.GetStudentCheckInLateHistory)

	r.POST("/checkin", api_device.EventCheckin)
	r.POST("/login", api_mobile.Login)

	return r, nil
}
