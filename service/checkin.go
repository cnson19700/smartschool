package service

import (
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

	device, err := repository.QueryDeviceByID(deviceID)
	if err != nil || device.RoomID == 0 {
		return "[Abnormal]: Device does not match any room"
	}

	schedule, err := repository.QueryScheduleByRoomTime(device.RoomID, checkinTime)
	if err != nil || schedule.ID == 0 {
		return "[Normal]: Time slot not in Schedule"
	}

	student, err := repository.QueryStudentBySID(studentID)
	if err != nil || student.ID == 0 {
		return "[Abnormal]: Student not recognize"
	}

	enrollment, _ := repository.QueryEnrollmentByStudentCourse(student.ID, schedule.CourseID)

	if enrollment.ID != 0 {

		checkAttend, _ := repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)

		if checkAttend.ID == 0 {
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

	device, err := repository.QueryDeviceByID(deviceID)
	if err != nil || device.RoomID == 0 {
		return "[Abnormal]: Device does not match any room"
	}

	course, err := repository.QueryCourseByID(courseID)
	if err != nil || course.ID == 0 {
		return "[Abnormal]: Course not exist"
	}

	schedule, err := repository.QueryScheduleByRoomTimeCourse(device.RoomID, checkinTime, course.ID)
	if err != nil || schedule.ID == 0 {
		return "[Normal]: Time slot not in Schedule"
	}

	student, err := repository.QueryStudentBySID(studentID)
	if err != nil || student.ID == 0 {
		return "[Abnormal]: Student not recognize"
	}

	enrollment, _ := repository.QueryEnrollmentByStudentCourse(student.ID, course.ID)

	if enrollment.ID != 0 {
		checkAttend, _ := repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)

		if checkAttend.ID == 0 {
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
