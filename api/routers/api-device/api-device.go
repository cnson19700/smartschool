package api_device

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api_students "github.com/smartschool/api/routers/api-mobile/api-students"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/service"
)

func EventCheckin(c *gin.Context) {
	var requestData dto.DeviceSignal
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		return
	}

	// fmt.Println(requestData.StudentId)
	// fmt.Println(requestData.Location)
	// fmt.Println(requestData.Timestamp)

	service.CheckIn(requestData)
}

func GetLateHistory(c *gin.Context) {
	id := c.Param("id")

	student, _ := api_students.GetStudentByID(c, id)

	listHistory := service.GetStudentHistoryFrom(student.ID, "Late")

	var historyElement = make([]dto.HistoryElement, 0)
	for i := 0; i < len(*listHistory); i++ {
		historyElement = append(historyElement, dto.HistoryElement{
			CourseName:    (*listHistory)[i].Scheduler.Course.CourseID + " - " + (*listHistory)[i].Scheduler.Course.Name,
			CheckinTime:   (*listHistory)[i].CheckInTime,
			CheckinStatus: (*listHistory)[i].CheckInStatus})
	}

	Mess := map[string]interface{}{
		"id":         student.ID,
		"student_id": student.StudentID,
		"name":       student.Name,
		"history":    historyElement,
	}

	c.JSON(http.StatusOK, Mess)

}
