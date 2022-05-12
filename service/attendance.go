package service

import (
	"errors"

	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
)

func GetFormRequestChangeAttendanceStatus(attendanceID uint) (*dto.RequestChangeAttendanceForm, []dto.TeacherList, error) {
	attendance, err := repository.QueryAttendanceByID(attendanceID)
	if err != nil {
		return nil, nil, err
	}

	teacherIDList, notFound, err := repository.QueryEnrollmentOfTeacherInCourse(attendance.Schedule.CourseID)
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
		ScheduleID:    attendance.ScheduleID,
		CourseName:    attendance.Schedule.Course.CourseID,
		Room:          attendance.Schedule.Room.RoomID,
		StartTime:     attendance.Schedule.StartTime,
		EndTime:       attendance.Schedule.EndTime,
		CurrentStatus: attendance.CheckInStatus,
	}

	return &schedule, teacherList, nil
}
