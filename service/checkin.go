package service

import (
	"fmt"
	"time"

	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

func CheckIn(deviceSignal dto.DeviceSignal) {

	status := ""
	entryTime := helper.ConvertDeviceTimestampToExact(deviceSignal.Timestamp)

	// if !helper.CheckValidDifferentTimeEntry(entryTime, 1*time.Minute) {
	// 	status = "[Abnormal]: Invalid Checkin time"
	// 	repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: deviceSignal.CardId, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})
	// 	return
	// }

	checkinType, checkinValue := helper.ClassifyCheckinCode(deviceSignal.CardId)

	switch checkinType {
	case "Card":
		status = recordCheckinCard(checkinValue, deviceSignal.CompanyTokenKey, entryTime)

	case "QR":
		status = recordCheckinQR(checkinValue, deviceSignal.CompanyTokenKey, entryTime)

	default:
		status = "[Abnormal]: Invalid format CardID"
	}

	repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: deviceSignal.CardId, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})
}

func recordCheckinCard(studentID string, deviceID string, checkinTime time.Time) string {
	fmt.Println(checkinTime)

	device := repository.QueryDeviceByID(deviceID)
	if device == nil {
		return "[Abnormal]: Device does not match any room"
	}

	schedule := repository.QueryScheduleByRoomTime(device.RoomID, checkinTime)
	if schedule == nil {
		return "[Normal]: Time slot not in Schedule"
	}

	student := repository.QueryStudentBySID(studentID)
	if student == nil {
		return "[Abnormal]: Student not recognize"
	}

	enrollment := repository.QueryEnrollmentByStudentCourse(student.ID, schedule.CourseID)

	if enrollment != nil {

		checkAttend := repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)

		if checkAttend == nil {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
				checkinStatus = "Late"
			}

			//database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			repository.CreateAttendance(entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			return "[Normal]: Checkin Success"
		} else {
			return "[Normal]: Checkin exist"
		}
	} else {
		return "[Normal]: Student dont take this course"
	}
}

func recordCheckinQR(checkinValues string, deviceID string, checkinTime time.Time) string {
	studentID, courseID := helper.ParseData(checkinValues)

	device := repository.QueryDeviceByID(deviceID)
	if device == nil {
		return "[Abnormal]: Device does not match any room"
	}

	course := repository.QueryCourseByID(courseID)
	if course == nil {
		return "[Abnormal]: Course not exist"
	}

	schedule := repository.QueryScheduleByRoomTimeCourse(device.RoomID, checkinTime, course.ID)
	if schedule == nil {
		return "[Normal]: Time slot not in Schedule"
	}

	student := repository.QueryStudentBySID(studentID)
	if student == nil {
		return "[Abnormal]: Student not recognize"
	}

	enrollment := repository.QueryEnrollmentByStudentCourse(student.ID, course.ID)

	if enrollment != nil {
		checkAttend := repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)

		if checkAttend == nil {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
				checkinStatus = "Late"
			}

			//database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			repository.CreateAttendance(entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			return "[Normal]: Checkin Success"
		} else {
			return "[Normal]: Checkin exist"
		}
	} else {
		return "[Normal]: Student dont take this course"
	}
}
