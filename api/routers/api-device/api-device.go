package api_device

import (
	//"net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/model/dto"

	"github.com/smartschool/service"
)

func EventCheckin(c *gin.Context) {
	var requestData dto.DeviceSignal
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		return
	}

	service.CheckIn(requestData)
}

func GetStudentCheckInLateHistory(c *gin.Context) {
	id := c.Param("id")

	studentFound, checkinHistoryList := service.GetCheckInHistoryBySID(id, "Late")

	if studentFound == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found!!!",
		})
		return
	}

	resp := map[string]interface{}{
		"id":         studentFound.ID,
		"student_id": studentFound.StudentID,
		"history":    checkinHistoryList,
		//"name":       studentFound.Name,
	}

	c.JSON(http.StatusOK, resp)

}

func GetCourseByID(c *gin.Context) {
	id := c.Param("id")

	course, err := service.GetCourseByCourseID(id)
	var Response map[string]interface{}
	if err != nil {
		Response = map[string]interface{}{
			"error": "Not Found",
		}
		c.JSON(http.StatusNotFound, Response)
	} else {

		Response = map[string]interface{}{
			"id":        course.ID,
			"course_id": course.CourseID,
			"name":      course.Name,
			"error":     nil,
		}

		c.JSON(http.StatusOK, Response)
	}

}

func GetCourses(c *gin.Context) {
	course, err := service.GetCourses()
	var Response map[string]interface{}
	if err != nil {
		Response = map[string]interface{}{
			"error": "Not Found",
		}
		c.JSON(http.StatusNotFound, Response)
	} else {

		Response = map[string]interface{}{
			"courses": course,
			"error":   nil,
		}

		c.JSON(http.StatusOK, Response)
	}

}

// func GetLateHistory(c *gin.Context) {
// 	id := c.Param("id")

// 	student := service.GetStudentByID(id)

// 	listHistory := service.GetStudentHistoryFrom(student.ID, "Late")

// 	var historyElement = make([]dto.HistoryElement, 0)
// 	for i := 0; i < len(*listHistory); i++ {
// 		historyElement = append(historyElement, dto.HistoryElement{
// 			CourseName:    (*listHistory)[i].Scheduler.Course.CourseID + " - " + (*listHistory)[i].Scheduler.Course.Name,
// 			CheckinTime:   (*listHistory)[i].CheckInTime,
// 			CheckinStatus: (*listHistory)[i].CheckInStatus})
// 	}
// 	// var stu model.Student
// 	// model.DbEntity.Where("id = ?", id).Find(&stu)

// 	Mess := map[string]interface{}{
// 		"id":         student.ID,
// 		"student_id": student.StudentID,
// 		"name":       student.Name,
// 		"history":    historyElement,
// 	}

// 	c.JSON(http.StatusOK, Mess)

// }
