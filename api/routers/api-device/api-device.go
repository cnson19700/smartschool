package api_device

import (
	"net/http"

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

func GetLateHistory(c *gin.Context) {
	id := c.Param("id")

	student := service.GetStudentByID(id)

	listHistory := service.GetStudentHistoryFrom(student.ID, "Late")

	var historyElement = make([]dto.HistoryElement, 0)
	for i := 0; i < len(*listHistory); i++ {
		historyElement = append(historyElement, dto.HistoryElement{CourseName: (*listHistory)[i].Scheduler.Course.CourseID + " - " + (*listHistory)[i].Scheduler.Course.Name, CheckinTime: (*listHistory)[i].CheckInTime, CheckinStatus: (*listHistory)[i].CheckInStatus})
	}
	// var stu model.Student
	// model.DbEntity.Where("id = ?", id).Find(&stu)

	Mess := map[string]interface{}{
		"id":         student.ID,
		"student_id": student.StudentID,
		"name":       student.Name,
		"history":    historyElement,
	}

	c.JSON(http.StatusOK, Mess)

}