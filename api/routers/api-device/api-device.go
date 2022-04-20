package api_device

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"

	"github.com/smartschool/service"
)

func EventCheckin(c *gin.Context) {
	var requestData dto.DeviceSignal
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	err = service.CheckIn(requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// func GetStudentCheckInLateHistory(c *gin.Context) {
// 	id := c.Param("id")

// 	studentFound, checkinHistoryList := service.GetCheckInHistoryBySID(id, "Late")

// 	if studentFound == nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Student not found!!!",
// 		})
// 		return
// 	}

// 	resp := map[string]interface{}{
// 		"id":         studentFound.ID,
// 		"student_id": studentFound.StudentID,
// 		"history":    checkinHistoryList,
// 		//"name":       studentFound.Name,
// 	}

// 	c.JSON(http.StatusOK, resp)

// }

// func GetCourseByID(c *gin.Context) {
// 	id := c.Param("id")

// 	course, err := service.GetCourseByCourseID(id)
// 	var Response map[string]interface{}
// 	if err != nil {
// 		Response = map[string]interface{}{
// 			"error": "Not Found",
// 		}
// 		c.JSON(http.StatusNotFound, Response)
// 	} else {

// 		Response = map[string]interface{}{
// 			"id":        course.ID,
// 			"course_id": course.CourseID,
// 			"name":      course.Name,
// 			"error":     nil,
// 		}

// 		c.JSON(http.StatusOK, Response)
// 	}

// }

func GetAllCourses(c *gin.Context) {
	course, err := repository.QueryAllCourses()
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
