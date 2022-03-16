package service

import (
	"fmt"
	"strings"
	"time"

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
		fmt.Println("Check in Card called")
		recordCheckinCard(checkinValue, deviceSignal.CompanyTokenKey, time.Unix(ConvertDeviceTimestampToExact(deviceSignal.Timestamp), 0))
		fmt.Println("Service checkincalled")

	case "QR":
		// recordCheckinQR(checkinValue, deviceSignal.CompanyTokenKey, time.Unix(ConvertDeviceTimestampToExact(deviceSignal.Timestamp), 0))
		fmt.Println("Service checkin QR called")

	default:
		return
	}
}

func recordCheckinCard(studentID string, deviceID string, checkinTime time.Time) {
	fmt.Println(checkinTime)

	device := repository.QueryDeviceByID(deviceID)
	if device == nil {
		fmt.Println("Device does not match any room!!!")
		return
	}

	schedule := repository.QueryScheduleByRoomTime(device.RoomID, checkinTime)
	if schedule == nil {
		fmt.Println("Time slot not in Schedule!!!")
		return
	}

	student := repository.QueryStudentBySID(studentID)
	if student == nil {
		fmt.Println("Student not recognize!!!")
		return
	}

	verify := repository.QueryEnrollmentByStudentCourse(student.ID, schedule.CourseID)

	if verify != nil {

		checkAttend := repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)

		if checkAttend == nil {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
				checkinStatus = "Late"
			}

			database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			fmt.Println("Checkin Success!!!")
		} else {
			fmt.Println("Checkin exist!!!")
		}
	} else {
		fmt.Println("Student dont take this course!!!")
	}
}

/*
func recordCheckinQR(checkinValues string, deviceID string, checkinTime time.Time) {
	studentID, courseID := parseData(checkinValues)

	device := repository.QueryDeviceByID(deviceID)
	if device.RoomID == 0 {
		fmt.Println("Device does not match any room!!!")
		return
	}

	course := repository.QueryCourseByID(courseID)
	if course.ID == 0 {
		fmt.Println("Course not exist!!!")
		return
	}

	schedule := repository.QueryScheduleByRoomTimeCourse(device.Room.RoomID, checkinTime, courseID)
	if schedule.ID == 0 {
		fmt.Println("Time slot not in Schedule!!!")
		return
	}

	student := repository.QueryStudentBySID(studentID)
	if student.ID == 0 {
		fmt.Println("Student not recognize!!!")
		return
	}

	verify := repository.QueryEnrollmentByStudentCourse(studentID, courseID)

	if verify.ID != 0 {
		checkAttend := repository.QueryAttendanceByStudentSchedule(studentID, schedule.ID)

		if checkAttend.ID == 0 {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
				checkinStatus = "Late"
			}

			database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			fmt.Println("Checkin Success!!!")
		} else {
			fmt.Println("Checkin exist!!!")
		}
	} else {
		fmt.Println("Student dont take this course!!!")
	}
}
*/

func ConvertDeviceTimestampToExact(timestamp int64) int64 {
	tempTime := time.Unix(timestamp, 0)
	tempTime = tempTime.Add((-1) * time.Hour * 7)
	// if tempTime.Unix() > time.Now().Unix() {
	// 	tempTime = time.Now()
	// }
	return tempTime.Unix()
}

func parseData(checkinValues string) (string, string) {
	res := strings.Split(checkinValues, "-")
	return res[0], res[1]
}
