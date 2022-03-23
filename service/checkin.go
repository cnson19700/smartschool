package service

import (
	"time"

	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)


const AcceptLate int = 20
const AcceptEarly int = 20

func CheckIn(deviceSignal dto.DeviceSignal) error {

	var status string = ""
	var err error = nil
	entryTime := helper.ConvertDeviceTimestampToExact(deviceSignal.Timestamp)

	// if !helper.CheckValidDifferentTimeEntry(entryTime, 1*time.Minute) {
	// 	status = "[Abnormal]: Invalid Checkin time"
	// 	repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: deviceSignal.CardId, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})
	// 	return
	// }

	checkinType, checkinValue, err := helper.ClassifyCheckinCode(deviceSignal.CardId)

	switch checkinType {
	case "Card":
		status, err = recordCheckinCard(checkinValue, deviceSignal.CompanyTokenKey, entryTime)

	case "QR":
		status = recordCheckinQR(checkinValue, deviceSignal.CompanyTokenKey, entryTime)

	default:
		status = "[Abnormal]: Invalid format CardID"
	}

	repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: deviceSignal.CardId, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})

	return err
}

func recordCheckinCard(studentID string, deviceID string, checkinTime time.Time) (string, error) {

	device, notFound, err := repository.QueryDeviceByID(deviceID)
	if err != nil {
		return "[Abnormal]: Error when query Device", err
	}
	if notFound {
		return "[Abnormal]: Device does not match any room", nil
	}

	student, notFound, err := repository.QueryStudentBySID(studentID)
	if err != nil {
		return "[Abnormal]: Error when query Student by SID", err
	}
	if notFound {
		return "[Abnormal]: Student not recognize", nil
	}

	schedule, notFound, err := repository.QueryScheduleByRoomTime(device.RoomID, checkinTime)
	if err != nil {
		return "[Abnormal]: Error when query Schedule", err
	}
	if notFound {
		return "[Normal]: Time slot not in any Schedule", nil
	}

	_, notFound, err = repository.QueryEnrollmentByStudentCourse(student.ID, schedule.CourseID)
	if err != nil {
		return "[Abnormal]: Error when query Student Course Enrollment", err
	}

	if !notFound {
		_, notFound, err = repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)
		if err != nil {
			return "[Abnormal]: Error when query Attendance", err
		}

		if notFound {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * time.Duration(AcceptLate)) {
				checkinStatus = "Late"
			}

			//database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			err = repository.CreateAttendance(entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			if err != nil {
				return "[Abnormal]: Error when create ttendance", err
			}

			return "[Normal]: Checkin Success", nil
		} else {
			return "[Normal]: Checkin exist", nil
		}
	} else {
		return "[Normal]: Student dont take this course", nil
	}
}

func recordCheckinQR(checkinValues string, deviceID string, checkinTime time.Time) string {
	// studentID, courseID := helper.ParseData(checkinValues)

	// device, err := repository.QueryDeviceByID(deviceID)
	// if err != nil || device.RoomID == 0 {
	// 	return "[Abnormal]: Device does not match any room"
	// }

	// course, err := repository.QueryCourseByID(courseID)
	// if err != nil || course.ID == 0 {
	// 	return "[Abnormal]: Course not exist"
	// }

	// schedule, err := repository.QueryScheduleByRoomTimeCourse(device.RoomID, checkinTime, course.ID)
	// if err != nil || schedule.ID == 0 {
	// 	return "[Normal]: Time slot not in Schedule"
	// }

	// student, err := repository.QueryStudentBySID(studentID)
	// if err != nil || student.ID == 0 {
	// 	return "[Abnormal]: Student not recognize"
	// }

	// enrollment, _ := repository.QueryEnrollmentByStudentCourse(student.ID, course.ID)

	// if enrollment.ID != 0 {
	// 	checkAttend, _ := repository.QueryAttendanceByStudentSchedule(student.ID, schedule.ID)

	// 	if checkAttend.ID == 0 {
	// 		checkinStatus := "Attend"
	// 		if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
	// 			checkinStatus = "Late"
	// 		}

	// 		//database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
	// 		repository.CreateAttendance(entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
	// 		return "[Normal]: Checkin Success"
	// 	} else {
	// 		return "[Normal]: Checkin exist"
	// 	}
	// } else {
	// 	return "[Normal]: Student dont take this course"
	// }

	return "Under maintainance"
}
