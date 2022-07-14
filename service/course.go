package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
)

func GetAttendanceInCourseOneUser(courseID uint, userID uint) ([]dto.AttendanceListElement, error) {

	currentTime := time.Now().UTC()
	scheduleList, notFound, err := repository.QueryListScheduleByCourse(courseID, currentTime)
	if err != nil || notFound {
		return nil, err
	}

	var scheduleIDList []uint
	//scheduleMap := make(map[uint]entity.Schedule)
	for i := 0; i < len(scheduleList); i++ {
		scheduleIDList = append(scheduleIDList, scheduleList[i].ID)
		//scheduleMap[scheduleList[i].ID] = scheduleList[i]
	}

	attendList, _, err := repository.QueryListAttendanceByUserSchedule(userID, scheduleIDList)
	if err != nil {
		return nil, err
	}

	var checkinTime *time.Time
	var checkinStatus string
	resultList := make([]dto.AttendanceListElement, 0)

	// for i := 0; i < len(attendList); i++ {
	// 	resultList = append(resultList, dto.AttendanceListElement{
	// 		StartTime:     scheduleMap[attendList[i].ScheduleID].StartTime,
	// 		EndTime:       scheduleMap[attendList[i].ScheduleID].EndTime,
	// 		CheckinTime:   attendList[i].CheckInTime,
	// 		Room:          scheduleMap[attendList[i].ScheduleID].Room.RoomID,
	// 		CheckinStatus: attendList[i].CheckInStatus,
	// 	})
	// }

	for i := 0; i < len(scheduleList); i++ {

		checkinTime = nil
		checkinStatus = ""
		for j := 0; j < len(attendList); j++ {
			if scheduleList[i].ID == attendList[j].ScheduleID {
				checkinTime = &attendList[j].CheckInTime
				checkinStatus = attendList[j].CheckInStatus
			}
		}

		if checkinStatus == "" && currentTime.Before(scheduleList[i].EndTime) && (currentTime.After(scheduleList[i].StartTime) || currentTime.Equal(scheduleList[i].StartTime)) {
			continue
		}

		resultList = append(resultList, dto.AttendanceListElement{
			StartTime:     scheduleList[i].StartTime,
			EndTime:       scheduleList[i].EndTime,
			Room:          scheduleList[i].Room.RoomID,
			CheckinTime:   checkinTime,
			CheckinStatus: checkinStatus,
		})
	}

	return resultList, nil
}

func GetCourseBasicInfoByID(id uint) (*dto.CourseReportPartElement, error) {
	course, notFound, err := repository.QueryCourseBasicInfoByID(id)
	if err != nil || notFound {
		return nil, err
	}

	return course, nil
}

func GetListCourseByUserSemester(userID uint, semesterID uint) ([]dto.CourseReportListElement, error) {
	courseIDInSemseterList, notFound, err := repository.QueryListCourseIDBySemester(semesterID)
	if err != nil || notFound {
		return nil, err
	}

	courseIDList, notFound, err := repository.QueryEnrollmentByListCourse(userID, courseIDInSemseterList)
	if err != nil || notFound {
		return nil, err
	}

	courseList, notFound, err := repository.QueryListCourseBasicInfoByID(courseIDList)
	if err != nil || notFound {
		return nil, err
	}

	currentTime := time.Now().UTC()
	resultList := make([]dto.CourseReportListElement, 0)
	var totalCount, attendCount int64
	var absenceCount uint
	var currentScheduleList []uint
	for i := 0; i < len(courseList); i++ {
		totalCount, err = repository.CountScheduleOfFullCourse(courseList[i].ID)
		if err != nil {
			continue
		}

		currentScheduleList, notFound, err = repository.QueryCurrentScheduleIDOfCourse(courseList[i].ID, currentTime)
		if err != nil {
			continue
		}
		attendCount = 0
		if !notFound {
			attendCount, err = repository.CountAttendanceOfSchedule(userID, currentScheduleList)
			if err != nil {
				continue
			}
		}

		absenceCount = uint(len(currentScheduleList) - int(attendCount))

		schedule_id, notFound, err := repository.QueryExistCurrentSchedule(courseList[i].ID, currentTime)
		if err == nil && !notFound {
			notFound, err = repository.ExistAttendanceByUserSchedule(userID, schedule_id)
			if err == nil && !notFound {
				attendCount += 1
			}
		}

		resultList = append(resultList, dto.CourseReportListElement{
			ID:         courseList[i].ID,
			CourseID:   courseList[i].CourseID,
			Name:       courseList[i].Name,
			Attendance: uint(attendCount),
			Absence:    absenceCount,
			Total:      uint(totalCount),
		})
	}

	return resultList, nil

}

func DeleteCourseBySemester(c *gin.Context) {
	w := c.Writer

	err := deleteCourseBySemester(1, 7)
	if err != nil {
		w.Write([]byte("Cannot delete course(s) in selected semester for user"))
		return
	}
}

func deleteCourseBySemester(facultyID uint, timezoneOffset float32) error {

	if timezoneOffset > 14 || timezoneOffset < -12 {
		return errors.New("invalid time zone, action delete course in semester is stopped")
	}

	currentDateTime := time.Now().UTC().Add(time.Hour * time.Duration(timezoneOffset))

	semesterID, notFound, err := repository.QuerySemesterByFacultyTime(facultyID, currentDateTime)
	if err != nil || notFound {
		return errors.New("invalid semester found, action delete course in semester is stopped")
	}

	listCourseIDInSemester, notFound, err := repository.QueryListCourseIDBySemester(semesterID)
	if err != nil {
		return err
	}
	if notFound {
		return nil
	}

	err = repository.DeleteTeacherCourseByListCourseID(listCourseIDInSemester)
	if err != nil {
		return err
	}

	err = repository.DeleteCourseByListCourseID(listCourseIDInSemester)
	if err != nil {
		return err
	}

	return nil
}
