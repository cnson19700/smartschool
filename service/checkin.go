package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

// c *gin.Context, deviceSignal dto.DeviceSignal

func CheckIn(deviceSignal dto.DeviceSignal) {
	checkinType, checkinValue := helper.ClassifyCheckinCode(deviceSignal.CardId)
	//checkinType := "Card"
	switch checkinType {
	case "Card":
		// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		recordCheckin(checkinValue, deviceSignal.CompanyTokenKey, time.Unix(ConvertDeviceTimestampToExact(deviceSignal.Timestamp), 0))

		//t0 := helper.StringToTimeUTC("2022-02-16T9:59:00Z")
		//fmt.Println(t0)
		//recordCheckin("100", "D1", t0)
		//var checkAttend entity.Attendance
		//database.DbInstance.Select("id").Where("student_id = ? AND course_id = ? AND room_id = ? AND end_time > ?", 1, 2, 1, t0.In(loc)).Find(&checkAttend)

		//fmt.Println(checkAttend.ID)
		fmt.Println("Service checkincalled")

	case "QR":
		recordCheckinQR(checkinValue, deviceSignal.CompanyTokenKey, time.Unix(ConvertDeviceTimestampToExact(deviceSignal.Timestamp), 0))

	default:
		return
	}
}

func recordCheckin(studentID string, deviceID string, checkinTime time.Time) {
	// fmt.Println(checkinTime)

	// var device entity.Device
	// database.DbInstance.Select("room_id").Where("device_id = ?", deviceID).Find(&device)
	// if device.RoomID == 0 {
	// 	fmt.Println("Device does not match any room!!!")
	// 	return
	// }

	// var scheduler entity.Scheduler
	// database.DbInstance.Order("end_time").Select("id", "course_id", "start_time", "end_time").Where("room_id = ? AND end_time >= ?", device.RoomID, checkinTime).Find(&scheduler)
	// if scheduler.ID == 0 {
	// 	fmt.Println("Time slot not in Schedule!!!")
	// 	return
	// }

	// var student entity.Student
	// database.DbInstance.Select("id").Where("student_id = ?", studentID).First(&student)
	// if student.ID == 0 {
	// 	fmt.Println("Student not recognize!!!")
	// 	return
	// }

	// var verify entity.StudentCourse
	// database.DbInstance.Select("id").Where("student_id = ? AND course_id = ?", student.ID, scheduler.CourseID).Find(&verify)

	// if verify.ID != 0 {
	// 	var checkAttend entity.Attendance
	// 	database.DbInstance.Select("id").Where("student_id = ? AND scheduler_id = ?", student.ID, scheduler.ID).Find(&checkAttend)

	// 	if checkAttend.ID == 0 {
	// 		checkinStatus := "Attend"
	// 		if timeDiff := checkinTime.Sub(scheduler.StartTime); timeDiff > (time.Minute * 20) {
	// 			checkinStatus = "Late"
	// 		}

	// 		database.DbInstance.Create(&entity.Attendance{StudentID: student.ID, SchedulerID: scheduler.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
	// 		fmt.Println("Checkin Success!!!")
	// 	} else {
	// 		fmt.Println("Checkin exist!!!")
	// 	}
	// } else {
	// 	fmt.Println("Student dont take this course!!!")
	// }
}

func recordCheckinQR(checkinValues string, deviceID string, checkinTime time.Time) {
	studentID, courseID := parseData(checkinValues)
	var c *gin.Context

	device := repository.QueryDeviceByID(deviceID)
	if device.RoomID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Device does not match any room!!!",
		})
		return
	}

	_, err := repository.QueryCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Course not exist!!!",
		})
		return
	}

	schedule, err := repository.QueryScheduleByRoomTimeCourse(device.Room.RoomID, checkinTime, courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Time slot not in Schedule!!!"})
		return
	}

	student, err := repository.QueryStudentBySID(studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not recognize!!!"})
		return
	}

	_, err = repository.QueryEnrollmentByStudentCourse(studentID, courseID)

	if err != nil {
		_, err := repository.QueryAttendanceByStudentSchedule(studentID, schedule.ID)

		if err != nil {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
				checkinStatus = "Late"
			}

			database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			c.JSON(http.StatusOK, gin.H{
				"message": "Checkin Success!!!"})
		} else {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": "Checkin exist!!!"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student dont take this course!!!"})
	}
}

func ConvertDeviceTimestampToExact(timestamp int64) int64 {
	tempTime := time.Unix(timestamp, 0)
	tempTime = tempTime.Add((-1) * time.Hour * 7)
	if tempTime.Unix() > time.Now().Unix() {
		tempTime = time.Now()
	}
	return tempTime.Unix()
}

func parseData(checkinValues string) (string, string) {
	res := strings.Split(checkinValues, "-")
	return res[0], res[1]
}
