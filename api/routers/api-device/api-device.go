package api_device

import (
	"github.com/gin-gonic/gin"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/service"
)

func EventCheckin(c *gin.Context) {
	var requestData dto.DeviceSignal
	if err := c.ShouldBindJSON(&requestData); err != nil {
		return
	}

	// fmt.Println(requestData.StudentId)
	// fmt.Println(requestData.Location)
	// fmt.Println(requestData.Timestamp)

	service.CheckIn(requestData)
}
