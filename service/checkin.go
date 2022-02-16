package service

import (
	"fmt"
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/entity"
	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
)

// c *gin.Context, deviceSignal dto.DeviceSignal

func CheckIn(deviceSignal dto.DeviceSignal) {
	checkinType, checkinValue := helper.ClassifyCheckinCode(deviceSignal.CardId)
	//checkinType := "Card"
	switch checkinType {
	case "Card":
		recordCheckin(checkinValue, deviceSignal.CompanyTokenKey, deviceSignal.TimeStamp.UTC())

		//t0 := helper.StringToTimeUTC("2022-02-16T9:59:00Z")
		//fmt.Println(t0)
		//recordCheckin("100", "D1", t0)
		//var checkAttend entity.Attendance
		//database.DbInstance.Select("id").Where("student_id = ? AND course_id = ? AND room_id = ? AND end_time > ?", 1, 2, 1, t0.In(loc)).Find(&checkAttend)

		//fmt.Println(checkAttend.ID)

	case "QR":
	default:
		return
	}
}

func recordCheckin(studentID string, deviceID string, checkinTime time.Time) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	fmt.Println(checkinTime.In(loc))

	var device entity.Device
	database.DbInstance.Select("room_id").Where("device_id = ?", deviceID).Find(&device)

	var scheduler entity.Scheduler
	database.DbInstance.Select("course_id", "end_time").Where("room_id = ? AND start_time <= ? AND end_time > ?", device.RoomID, checkinTime, checkinTime).Find(&scheduler)

	if scheduler.CourseID == 0 {
		fmt.Println("Schedule not set!!!")
		return
	}

	var student entity.Student
	database.DbInstance.Select("id").Where("student_id = ?", studentID).First(&student)

	var verify entity.StudentCourse
	database.DbInstance.Select("id").Where("student_id = ? AND course_id = ?", student.ID, scheduler.CourseID).Find(&verify)

	if verify.ID != 0 {
		var checkAttend entity.Attendance
		database.DbInstance.Select("id").Where("student_id = ? AND course_id = ? AND room_id = ? AND end_time > ?", student.ID, scheduler.CourseID, device.RoomID, checkinTime).Find(&checkAttend)

		if checkAttend.ID == 0 {
			database.DbInstance.Create(&entity.Attendance{StudentID: student.ID, CourseID: scheduler.CourseID, RoomID: device.RoomID, CheckInTime: checkinTime, EndTime: scheduler.EndTime, CheckInStatus: "Attend"})
			fmt.Println("Checkin Success!!!")
		} else {
			fmt.Println("Checkin exist!!!")
		}
	} else {
		fmt.Println("Record not found")
	}
}
