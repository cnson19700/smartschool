package service

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/smartschool/helper"
	"github.com/smartschool/lib/constant"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
	"github.com/smartschool/service/fireapp"
	"golang.org/x/crypto/bcrypt"
)

func CheckIn(deviceSignal dto.DeviceSignal) error {

	var status string = ""
	var checkinValue string = deviceSignal.CardId
	var err error = nil
	entryTime := helper.ConvertDeviceTimestampToExact(deviceSignal.Timestamp)

	if !helper.CheckValidDifferentTimeEntry(entryTime, constant.AcceptDeviceSignalDelay) {
		status = constant.CheckinStatus_InvalidCheckinTime
		repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: checkinValue, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})
		return nil
	}

	checkinType, err := helper.ClassifyCheckinCode(deviceSignal.CardId)

	switch checkinType {
	case constant.CheckinType_Card:

		userId, notFound, errFind := findRequestUserByCardID(checkinValue)
		if errFind != nil || notFound {
			status = constant.CheckinStatus_InvalidCardUserNotFound
		} else {
			checkinValue = strconv.Itoa(int(userId))
			userRole, errFind := findUserRole(userId)
			if errFind != nil {
				status = constant.CheckinStatus_InvalidCardRoleNotFound
			} else {
				status, err = recordCheckin(userId, userRole, deviceSignal.CompanyTokenKey, entryTime)
			}
		}

	case constant.CheckinType_QR:
		userId, isFormatCorrect, errParse := helper.ParseQR(checkinValue, entryTime)
		if !isFormatCorrect || errParse != nil {
			status = constant.CheckinStatus_InvalidQR
		} else {
			checkinValue = strconv.Itoa(int(userId))
			userRole, errFind := findUserRole(userId)
			if errFind != nil {
				status = constant.CheckinStatus_InvalidQRRoleNotFound
			} else {
				status, err = recordCheckin(userId, userRole, deviceSignal.CompanyTokenKey, entryTime)
			}
		}

	default:
		status = constant.CheckinStatus_InvalidInfo
	}

	repository.CreateLogCheckIn(entity.DeviceSignalLog{CardId: checkinValue, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: entryTime})

	return err
}

func recordCheckin(userID uint, userRole uint, deviceID string, checkinTime time.Time) (string, error) {

	roomId, notFound, err := repository.QueryRoomByDeviceID(deviceID)
	if err != nil {
		NotiFail(userID, constant.CheckinStatus_ErrorQueryDevice)
		return constant.CheckinStatus_ErrorQueryDevice, err
	}
	if notFound {
		NotiFail(userID, constant.CheckinStatus_DeviceNotFound)
		return constant.CheckinStatus_DeviceNotFound, nil
	}

	var isScheduleForeseen bool = false
	schedule, notFound, err := repository.QueryScheduleByRoomTime(roomId, checkinTime)
	if err != nil {
		NotiFail(userID, constant.CheckinStatus_ErrorQuerySchedule)
		return constant.CheckinStatus_ErrorQuerySchedule, err
	}
	needCheckNextSchedule := notFound

	for !isScheduleForeseen {
		if needCheckNextSchedule {
			temp := schedule
			schedule, notFound, err = repository.QueryScheduleByRoomTime(roomId, checkinTime.Add(constant.AcceptEarlyMinute))
			isScheduleForeseen = true
			if err != nil {
				NotiFail(userID, constant.CheckinStatus_ErrorQuerySchedule)
				return constant.CheckinStatus_ErrorQuerySchedule, err
			}
			if notFound {
				NotiFail(userID, constant.CheckinStatus_ScheduleNotFound)
				return constant.CheckinStatus_ScheduleNotFound, nil
			}
			if schedule.ID == temp.ID {
				NotiFail(userID, constant.CheckinStatus_SameScheduleSpam)
				return constant.CheckinStatus_SameScheduleSpam, nil
			}
		}

		if userRole == constant.StudentRole {
			notFound, err = repository.ExistEnrollmentByStudentCourse(userID, schedule.CourseID)
		} else if userRole == constant.TeacherRole {
			notFound, err = repository.ExistEnrollmentByTeacherCourse(userID, schedule.CourseID)
		} else {
			NotiFail(userID, constant.CheckinStatus_AmbiguousUserRole)
			return constant.CheckinStatus_AmbiguousUserRole, err
		}
		if err != nil {
			NotiFail(userID, constant.CheckinStatus_ErrorQueryEnrollment)
			return constant.CheckinStatus_ErrorQueryEnrollment, err
		}

		if !notFound {
			notFound, err = repository.ExistAttendanceByUserSchedule(userID, schedule.ID)
			if err != nil {
				NotiFail(userID, constant.CheckinStatus_ErrorQueryAttendance)
				return constant.CheckinStatus_ErrorQueryAttendance, err
			}

			if notFound {
				checkinStatus := constant.CheckinStatus_Attend
				if timeDiff := checkinTime.Sub(schedule.StartTime); timeDiff > constant.AcceptLateMinute {
					checkinStatus = constant.CheckinStatus_Late
				}

				// teacherId, err := repository.QueryTeacherIDByCourseID(schedule.CourseID)
				// if err != nil {
				// 	return "[Abnormal]: Error when query teacher in course", err
				// }

				err = repository.CreateAttendance(entity.Attendance{UserID: userID, ScheduleID: schedule.ID, TeacherID: 0, CheckInTime: checkinTime, CheckInStatus: checkinStatus})
				if err != nil {
					NotiFail(userID, constant.CheckinStatus_ErrorCreateAttendance)
					return constant.CheckinStatus_ErrorCreateAttendance, err
				}
				student, _ := repository.QueryStudentByID(fmt.Sprint(userID))
				MessageToNotify(student, schedule, checkinTime, checkinStatus)
				return constant.CheckinStatus_Success, nil
			} else if isScheduleForeseen {
				return constant.CheckinStatus_Exist, nil
			}

		} else if isScheduleForeseen {
			NotiFail(userID, constant.CheckinStatus_EnrollmentNotFound)
			return constant.CheckinStatus_EnrollmentNotFound, nil
		}

		needCheckNextSchedule = true
	}
	NotiFail(userID, constant.CheckinStatus_ErrorLogic)
	return constant.CheckinStatus_ErrorLogic, nil
}

func GenerateQREncodeString(userId uint) (string, error) {
	currentDateTime, _ := time.Now().UTC().MarshalText()

	hashedSecretKeyByte, bcryptError := bcrypt.GenerateFromPassword([]byte(constant.QRSecretKey+string(currentDateTime)), bcrypt.DefaultCost)
	if bcryptError != nil {
		return "", bcryptError
	}

	rawString := strconv.Itoa(int(userId)) + "|" + string(currentDateTime) + "|" + string(hashedSecretKeyByte)

	encodeString := base64.StdEncoding.EncodeToString([]byte(rawString))

	QRString := constant.QRPrefix + ":" + encodeString + "="

	return QRString, nil
}

func findUserRole(userId uint) (uint, error) {
	roleId, err := repository.QueryUserRoleIDByUserID(userId)
	if err != nil {
		return 0, err
	}

	return roleId, nil
}

func findRequestUserByCardID(code string) (uint, bool, error) {
	userId, notFound, err := repository.QueryUserByCardID(code)

	return userId, notFound, err
}
func MessageToNotify(student *entity.Student, schedule *entity.Schedule, checkinTime time.Time, checkinStatus string) {
	course_name, _, err1 := repository.QueryCourseBasicInfoByID(schedule.CourseID)
	room, _, err2 := repository.QueryRoomInfo(schedule.RoomID)
	msg := ""
	if err1 != nil {
		msg += "Fail in finding Course Name"
	}

	if err2 != nil {
		msg += "\n Fail in finding Room"
	}
	if msg == "" {
		msg = "Success"
	}
	data := map[string]string{
		"message":       msg,
		"course":        course_name.Name,
		"room":          room.RoomID,
		"shift":         schedule.StartTime.Format("2006-01-02 15:04:05") + "-" + schedule.EndTime.Format("2006-01-02 15:04:05"),
		"checkintime":   checkinTime.Format("2006-01-02 15:04:05"),
		"checkinstatus": checkinStatus,
	}
	fireapp.SendNotification(student.ID, data)
}

func NotiFail(studentID uint, msg string) {
	data := map[string]string{
		"message": "Fail to Checkin",
	}
	fireapp.SendNotification(studentID, data)
}
