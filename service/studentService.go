package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/entity"
	"github.com/smartschool/model/dto"
)

func getStudentByID(id string) *entity.Student {
	var student entity.Student
	database.DbInstance.Where("student_id = ?", id).First(&student)
	if student.ID == 0 {
		return nil
	}

	return &student
}

// func getStudentHistoryWithIdAndStatus(id int, status string) *[]entity.Attendance {
// 	var stat []entity.Attendance
// 	//DbInstance.Preload("Scheduler", "schedulers.course_id = ?", 2).Where("student_id = ? AND check_in_status = ?", 1, "Late").Find(&stat)
// 	database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat)

// 	// var result []entity.Attendance
// 	// for i := 0; i < len(stat); i++ {
// 	// 	result = append(result, entity.Attendance{ID: stat[i].ID, StudentID: stat[i].StudentID, CheckInTime: stat[i].CheckInTime, CheckInStatus: stat[i].CheckInStatus})
// 	// }
// 	result := append([]entity.Attendance{}, stat...)

// 	return &result
// }

func GetCheckInHistoryByStudentIDAndStatus(c *gin.Context, id string, status string) {

	student := getStudentByID(id)
	if student.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Student not recognize!!!",
		})
		return
	}

	var stat []entity.Attendance
	database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat)

	listHistory := append([]entity.Attendance{}, stat...)

	var historyElement = make([]dto.CheckInHistoryElement, 0)
	for i := 0; i < len(listHistory); i++ {
		historyElement = append(historyElement, dto.CheckInHistoryElement{
			CourseName:    listHistory[i].Scheduler.Course.CourseID + " - " + listHistory[i].Scheduler.Course.Name,
			CheckinTime:   listHistory[i].CheckInTime,
			CheckinStatus: listHistory[i].CheckInStatus})
	}

	Mess := map[string]interface{}{
		"id":         student.ID,
		"student_id": student.StudentID,
		"name":       student.Name,
		"history":    historyElement,
	}

	c.JSON(http.StatusOK, Mess)
}
