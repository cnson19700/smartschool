package routers

import (
	"github.com/gin-gonic/gin"
	api_device "github.com/smartschool/api/routers/api-device"
)

func Initialize() *gin.Engine {
	r := gin.New()

	r.GET("/late/:id", api_device.GetLateHistory)
	r.POST("/checkin", api_device.EventCheckin)

	return r
}
