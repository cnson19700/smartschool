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

	//mobileGroup.GET("/late/:id", api_device.GetStudentCheckInLateHistory)
	mobileGroup.PUT("/change-password", api_mobile.UpdatePassword)

	mobileUser := r.Group("/user")
	mobileUser.Use(authMw.GetAuthFunc())
	mobileUser.GET("/me", api_mobile.GetMe)
	mobileUser.GET("/course-attendance", api_mobile.GetCourseAttendanceOfOneUser)
	mobileUser.GET("/inday-attendance", api_mobile.GetInDayAttendance)
	mobileUser.GET("/get-qr", api_mobile.GetQREncodeString)
	mobileUser.POST("/update-notification-token", api_mobile.UpdateNotificationToken)
	mobileUser.GET("/courses-in-semester", api_mobile.GetCourseInSemesterOfOneUser)
	mobileUser.GET("/semesters", api_mobile.GetSemesterInFaculty)
	mobileUser.POST("/change-password-firsttime", api_mobile.ChangePasswordFirstTime)

	mobileUser.GET("/test-notification", api_mobile.TestNotification)

	r.POST("/checkin", api_device.EventCheckin)
	r.POST("/login", api_mobile.Login)
	//r.GET("/courses/:id", api_device.GetCourseByID)
	// r.GET("/courses", api_device.GetCourses)
	// r.GET("/teachercourses", api_web.GetTeacherCourse)

	return r, nil
}
