package service

import (
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

func GetAttendanceInCourseOneUser(courseID uint, userID uint) ([]dto.AttendanceListElement, error) {
	scheduleList, notFound, err := repository.QueryListScheduleByCourse(courseID)

	if err != nil || notFound {
		return nil, err
	}

	var scheduleIDList []uint
	scheduleMap := make(map[uint]entity.Schedule)
	for i := 0; i < len(scheduleList); i++ {
		scheduleIDList = append(scheduleIDList, scheduleList[i].ID)
		scheduleMap[scheduleList[i].ID] = scheduleList[i]
	}

	attendList, notFound, err := repository.QueryListAttendanceByUserSchedule(userID, scheduleIDList)
	if err != nil || notFound {
		return nil, err
	}

	resultList := make([]dto.AttendanceListElement, 0)
	for i := 0; i < len(attendList); i++ {
		resultList = append(resultList, dto.AttendanceListElement{
			StartTime:     scheduleMap[attendList[i].ScheduleID].StartTime,
			EndTime:       scheduleMap[attendList[i].ScheduleID].EndTime,
			CheckinTime:   attendList[i].CheckInTime,
			Room:          scheduleMap[attendList[i].ScheduleID].Room.RoomID,
			CheckinStatus: attendList[i].CheckInStatus,
		})
	}

	return resultList, nil
}

func GetCourseInfoByID(id uint) (*entity.Course, error) {
	course, notFound, err := repository.QueryCourseInfoByID(id)
	if err != nil || notFound {
		return nil, err
	}

	return course, nil
}
