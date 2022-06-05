package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
)

func GetAttendanceInCourseOneUser(courseID uint, userID uint) ([]dto.AttendanceListElement, error) {
	scheduleList, notFound, err := repository.QueryListScheduleByCourse(courseID, time.Now().UTC())
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

	resultList := make([]dto.CourseReportListElement, 0)
	var totalCount, attendCount int64
	var absenceCount uint
	var currentScheduleList []uint
	for i := 0; i < len(courseList); i++ {
		totalCount, err = repository.CountScheduleOfFullCourse(courseList[i].ID)
		if err != nil {
			continue
		}

		currentScheduleList, notFound, err = repository.QueryCurrentScheduleIDOfCourse(courseList[i].ID, time.Now().UTC())
		if err != nil {
			continue
		}
		attendCount = 0
		if !notFound {
			attendCount, err = repository.CountAttendanceOfSchedule(userID, currentScheduleList)
		}

		absenceCount = uint(len(currentScheduleList) - int(attendCount))
		if err != nil {
			continue
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
	request := struct {
		SemesterID uint `form:"semester_id" binding:"required"`
	}{}

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot capture data for this request",
		})
		return
	}

	err = deleteCourseBySemester(request.SemesterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messgae": "Cannot delete course(s) in selected semester for user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messgae": "success",
	})
}

func deleteCourseBySemester(semester_id uint) error {
	listCourseIDInSemester, notFound, err := repository.QueryListCourseIDBySemester(semester_id)
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