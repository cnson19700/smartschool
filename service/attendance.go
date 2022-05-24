package service

import (
	"errors"
	"time"

	"github.com/smartschool/lib/constant"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

func GetFormRequestChangeAttendanceStatus(userID uint, scheduleID uint) (*dto.RequestChangeAttendanceForm, []dto.TeacherList, error) {

	var checkinTime *time.Time = nil

	scheduleInfo, err := repository.QueryScheduleByID(scheduleID)
	if err != nil {
		return nil, nil, err
	}

	notFound, err := repository.ExistEnrollmentByStudentCourse(userID, scheduleInfo.CourseID)
	if err != nil || notFound {
		return nil, nil, errors.New("student does not enroll this course")
	}

	attendance, notFound, err := repository.QueryAttendanceStatusByUserSchedule(userID, scheduleID)
	if err != nil {
		return nil, nil, err
	}

	if notFound {
		if time.Since(scheduleInfo.EndTime) >= 0 {
			attendance.CheckInStatus = constant.Absence
		} else {
			attendance.CheckInStatus = constant.Unknown
		}
	} else {
		checkinTime = &attendance.CheckInTime
	}

	teacherIDList, notFound, err := repository.QueryEnrollmentOfTeacherInCourse(scheduleInfo.CourseID)
	if err != nil || notFound {
		return nil, nil, errors.New("teacher for this course not found")
	}

	teacherInfoList, notFound, err := repository.QueryListUserNameInfo(teacherIDList)
	if err != nil || notFound {
		return nil, nil, errors.New("teacher for this course not found")
	}

	teacherList := make([]dto.TeacherList, 0)
	for i := 0; i < len(teacherInfoList); i++ {
		teacherList = append(teacherList, dto.TeacherList{
			ID:   teacherInfoList[i].ID,
			Name: teacherInfoList[i].LastName + " " + teacherInfoList[i].FirstName,
		})
	}

	schedule := dto.RequestChangeAttendanceForm{
		ScheduleID:    scheduleInfo.ID,
		CourseName:    scheduleInfo.Course.CourseID,
		Room:          scheduleInfo.Room.RoomID,
		StartTime:     scheduleInfo.StartTime,
		EndTime:       scheduleInfo.EndTime,
		CheckInTime:   checkinTime,
		CurrentStatus: attendance.CheckInStatus,
	}

	return &schedule, teacherList, nil
}

func RequestChangeAttendanceStatus(userId uint, request dto.ChangeAttendanceStatusRequest) error {
	schedule, err := repository.QueryScheduleCourseSemesterByID(request.ScheduleID)
	if err != nil {
		return err
	}

	if time.Since(schedule.Course.Semester.EndTime) >= 0 {
		return errors.New("can not resolve requests of passed semester")
	}

	formIDList, err := repository.QueryListFormIDByUser(userId, request.ToUserID, schedule.Course.SemesterID)
	if err != nil {
		return err
	}
	if len(formIDList) > 0 {
		notFound, err := repository.ExistFormSchedule(formIDList, request.ScheduleID)
		if err != nil || !notFound {
			return errors.New("form already exists")
		}
	}

	notFound, err := repository.ExistEnrollmentByStudentCourse(userId, schedule.CourseID)
	if err != nil || notFound {
		return errors.New("student does not enroll this course")
	}

	notFound, err = repository.ExistEnrollmentByTeacherCourse(request.ToUserID, schedule.CourseID)
	if err != nil || notFound {
		return errors.New("student does not enroll this course")
	}

	recordRequest := entity.AttendanceForm{
		RequestUserID:        userId,
		ReceiveUserID:        request.ToUserID,
		SemesterID:           schedule.Course.SemesterID,
		RequestCheckInStatus: request.RequestCheckInStatus,
		Reason:               request.Reason,
		Schedules:            []entity.Schedule{{ID: request.ScheduleID}},
		//	ScheduleID:           request.ScheduleID,
		//	AttendanceID:         attendanceID,
		//	CheckInStatus:        oldStatus,
	}

	err = repository.CreateChangeAttendanceRequest(recordRequest)
	return err
}
