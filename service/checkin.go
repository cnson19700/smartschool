package service

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/smartschool/helper"
	"github.com/smartschool/lib/constant"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
	"golang.org/x/crypto/bcrypt"
)

func CheckIn(deviceSignal dto.DeviceSignal) error {

	var status string = ""
	var checkinValue string = deviceSignal.CardId
	var err error = nil
	entryTime := helper.ConvertDeviceTimestampToExact(deviceSignal.Timestamp)

	if !helper.CheckValidDifferentTimeEntry(entryTime, constant.AcceptDeviceSignalDelay) {
		status = "[Abnormal]: Invalid Checkin time - Delay over 5 second"
		repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: checkinValue, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})
		return nil
	}

	checkinType, err := helper.ClassifyCheckinCode(deviceSignal.CardId)

	switch checkinType {
	case "Card":

		userId, notFound, errFind := findRequestUser(checkinValue)
		if errFind != nil || notFound {
			status = "[Abnormal]: Invalid Card - User not found"
		} else {
			checkinValue = strconv.Itoa(int(userId))
			userRole, errFind := findUserRole(userId)
			if errFind != nil {
				status = "[Abnormal]: Invalid User Card - Role not found"
			} else {
				status, err = recordCheckin(userId, userRole, deviceSignal.CompanyTokenKey, entryTime)
			}
		}
		//status, err = recordCheckinCard(checkinValue, deviceSignal.CompanyTokenKey, entryTime)

	case "QR":
		userId, isFormatCorrect, errParse := helper.ParseQR(checkinValue, entryTime)
		if !isFormatCorrect || errParse != nil {
			status = "[Abnormal]: Invalid format QR or Expired QR"
		} else {
			checkinValue = strconv.Itoa(int(userId))
			userRole, errFind := findUserRole(userId)
			if errFind != nil {
				status = "[Abnormal]: Invalid User QR - Role not found"
			} else {
				status, err = recordCheckin(userId, userRole, deviceSignal.CompanyTokenKey, entryTime)
			}
		}

		//status, err = recordCheckinCard(checkinValue, deviceSignal.CompanyTokenKey, entryTime)
		//status = recordCheckinQR(checkinValue, deviceSignal.CompanyTokenKey, entryTime)
	default:
		status = "[Abnormal]: Invalid format CardID"
	}

	repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: checkinValue, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})

	return err
}

func recordCheckin(userID uint, userRole string, deviceID string, checkinTime time.Time) (string, error) {

	roomId, notFound, err := repository.QueryRoomByDeviceID(deviceID)
	if err != nil {
		return "[Abnormal]: Error when query Device", err
	}
	if notFound {
		return "[Abnormal]: Device does not match any room", nil
	}

	// student, notFound, err := repository.QueryStudentBySID(studentID)
	// if err != nil {
	// 	return "[Abnormal]: Error when query Student by SID", err
	// }
	// if notFound {
	// 	return "[Abnormal]: Student not recognize", nil
	// }

	var isScheduleForeseen bool = false
	schedule, notFound, err := repository.QueryScheduleByRoomTime(roomId, checkinTime)
	if err != nil {
		return "[Abnormal]: Error when query Schedule", err
	}
	needCheckNextSchedule := notFound

	for !isScheduleForeseen {
		if needCheckNextSchedule {
			temp := schedule
			schedule, notFound, err = repository.QueryScheduleByRoomTime(roomId, checkinTime.Add(constant.AcceptEarlyMinute))
			isScheduleForeseen = true
			if err != nil {
				return "[Abnormal]: Error when query Schedule", err
			}
			if notFound {
				return "[Normal]: Forseen time slot not in any Schedule", nil
			}
			if schedule.ID == temp.ID {
				return "[Normal]: Spam check-in", nil
			}
		}

		if userRole == constant.StudentRole {
			notFound, err = repository.ExistQueryEnrollmentByStudentCourse(userID, schedule.CourseID)
		} else if userRole == constant.TeacherRole {
			notFound, err = repository.ExistQueryEnrollmentByTeacherCourse(userID, schedule.CourseID)
		} else {
			return "[Abnormal]: Ambiguous user role", err
		}
		if err != nil {
			return "[Abnormal]: Error when query Student Course Enrollment", err
		}

		if !notFound {
			notFound, err = repository.ExistQueryAttendanceByUserSchedule(userID, schedule.ID)
			if err != nil {
				return "[Abnormal]: Error when query Attendance", err
			}

			if notFound {
				checkinStatus := "Attend"
				if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > constant.AcceptLateMinute {
					checkinStatus = "Late"
				}

				// teacherId, err := repository.QueryTeacherIDByCourseID(schedule.CourseID)
				// if err != nil {
				// 	return "[Abnormal]: Error when query teacher in course", err
				// }

				//database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
				err = repository.CreateAttendance(entity.Attendance{UserID: userID, ScheduleID: schedule.ID, TeacherID: 0, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
				if err != nil {
					return "[Abnormal]: Error when create Attendance", err
				}

				return "[Normal]: Checkin Success", nil
			} else if isScheduleForeseen {
				return "[Normal]: Checkin exist", nil
			}

		} else if isScheduleForeseen {
			return "[Normal]: Student dont take this course", nil
		}

		needCheckNextSchedule = true
	}

	return "[Abnormal]: Error in Logic Check-in", nil
}

/* func recordCheckinQR(checkinValues string, deviceID string, checkinTime time.Time) (string, error) {
	studentID, courseID := helper.ParseData(checkinValues)

	device, notFound, err := repository.QueryDeviceByID(deviceID)
	if err != nil {
		return "[Abnormal]: Error when query Device", err
	}
	if notFound {
		return "[Abnormal]: Device does not match any room", nil
	}

	course, notFound, err := repository.QueryCourseByID(courseID)
	if err != nil {
		return "[Abnormal]: Error when query Device", err
	}
	if notFound {
		return "[Abnormal]: Course does not exist", nil
	}

	schedule, notFound, err := repository.QueryScheduleByRoomTimeCourse(device.RoomID, checkinTime, course.ID)
	if err != nil {
		return "[Abnormal]: Error when query Schedule", err
	}
	if notFound {
		return "[Normal]: Forseen time slot not in any Schedule", nil
	}

	student, notFound, err := repository.QueryStudentBySID(studentID)
	if err != nil {
		return "[Abnormal]: Error when query Student by SID", err
	}
	if notFound {
		return "[Abnormal]: Student not recognize", nil
	}

	_, notFound, _ = repository.QueryEnrollmentByStudentCourse(student.ID, course.ID)
	if err != nil {
		return "[Abnormal]: Error when query Student Course Enrollment", err
	}

	if !notFound {
		_, err = repository.ExistQueryAttendanceByStudentSchedule(student.ID, schedule.ID)
		if err != nil {
			return "[Abnormal]: Error when query Attendance", err
		}

		if notFound {
			checkinStatus := "Attend"
			if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > (time.Minute * 20) {
				checkinStatus = "Late"
			}

			//database.DbInstance.Create(&entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			err = repository.CreateAttendance(entity.Attendance{UserID: student.ID, ScheduleID: schedule.ID, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
			if err != nil {
				return "[Abnormal]: Error when create Attendance", err
			}

			return "[Normal]: Checkin Success", nil
		} else {
			return "[Normal]: Checkin exist", nil
		}
	} else {
		return "[Normal]: Student dont take this course", nil
	}
}*/

func GenerateQREncodeString(userId uint) (string, error) {
	currentDateTime, _ := time.Now().UTC().MarshalText()

	hashedSecretKeyByte, bcryptError := bcrypt.GenerateFromPassword([]byte(constant.QRSecretKey + string(currentDateTime)), bcrypt.DefaultCost)
	if bcryptError != nil {
		return "", bcryptError
	}

	// student, err := repository.QueryStudentByID(fmt.Sprint(userId))
	// if err != nil {
	// 	return "", err
	// }

	rawString := strconv.Itoa(int(userId)) + "|" + string(currentDateTime) + "|" + string(hashedSecretKeyByte)

	encodeString := base64.StdEncoding.EncodeToString([]byte(rawString))

	QRString := constant.QRPrefix + ":" + encodeString + "="

	return QRString, nil
}

func findUserRole(userId uint) (string, error) {
	roleId, err := repository.QueryUserRoleIDByUserID(userId)
	if err != nil {
		return "", err
	}

	roleTitle, err := repository.QueryUserRoleDetailByRoleID(roleId)
	if err != nil {
		return "", err
	}

	return roleTitle, nil
}

func findRequestUserByCardID(code string) (uint, bool, error) {
	userId, notFound, err := repository.QueryUserByCardID(code)

	return userId, notFound, err
}
