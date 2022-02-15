package api_device

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartschool/model/dto"
)

func EventCheckin(c *gin.Context) {
	var requestData dto.EventCheckin
	if err := c.ShouldBindJSON(&requestData); err != nil {
		return
	}

	fmt.Println(requestData.StudentId)
	fmt.Println(requestData.Location)
	fmt.Println(requestData.Timestamp)

	return
}
