package service

import (
	"time"

	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
)

// func getStudentHistoryWithIdAndStatus(id int, status string) *[]entity.Attendance {
// 	var stat []entity.Attendance
// 	//DbInstance.Preload("Scheduler", "schedulers.course_id = ?", 2).Where("student_id = ? AND check_in_status = ?", 1, "Late").Find(&stat)
// 	database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat)

// 	// var result []entity.Attendance
// 	// for i := 0; i < len(stat); i++ {
// 	// 	result = append(result, entity.Attendance{ID: stat[i].ID, StudentID: stat[i].StudentID, CheckInTime: stat[i].CheckInTime, CheckInStatus: stat[i].CheckInStatus})
// 	// }
// 	result := append([]entity.Attendance{}, stat...)

// 	return &result
// }

// func GetCheckInHistoryBySID(sid string, status string) (*entity.Student, []dto.CheckInHistoryElement) {

// 	student, err := repo.QueryStudentBySID(sid)
// 	if err != nil {
// 		return nil, nil
// 	}

// 	listHistory := repo.QueryCheckinHistoryWithSIdAndStatus(student.ID, status)

// 	var historyElements = make([]dto.CheckInHistoryElement, 0)
// 	for i := 0; i < len(listHistory); i++ {
// 		historyElements = append(historyElements, dto.CheckInHistoryElement{
// 			CourseName:    listHistory[i].Scheduler.Course.CourseID + " - " + listHistory[i].Scheduler.Course.Name,
// 			CheckinTime:   listHistory[i].CheckInTime,
// 			CheckinStatus: listHistory[i].CheckInStatus})
// 	}

// 	return student, historyElements
// 	return nil, nil
// }

func GetMe(id string) (*dto.StudentProfile, error) {
	student, err := repository.QueryStudentByID(id)
	user := repository.QueryUserBySID(id)
	if err != nil {
		return &dto.StudentProfile{}, err
	}
	var genderStudent string
	if user.Gender == 0 {
		genderStudent = "male"
	} else if user.Gender == 1 {
		genderStudent = "female"
	}

	StudentProfile := dto.StudentProfile{
		StudentID:   student.StudentID,
		Name:        user.LastName + " " + user.FirstName,
		Class:       student.Batch,
		Email:       user.Email,
		Gender:      genderStudent,
		PhoneNumber: user.PhoneNumber,
	}
	return &StudentProfile, nil
}

func GetCheckInHistoryInDay(userID uint, facultyID uint, timezoneOffset int) ([]dto.CheckInHistoryListElement, error) {
	if timezoneOffset > 14 || timezoneOffset < -12 {
		return nil, nil
	}

	currentDateTime := time.Now().UTC().Add(time.Hour * time.Duration(timezoneOffset))
	//currentDateTime := time.Date(2022, 1, 12, 15, 0, 0, 0, time.UTC).Add(time.Hour * time.Duration(timezoneOffset))
	year, month, day := currentDateTime.Date()

	startDateTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	startDateTime = startDateTime.Add(time.Hour * time.Duration(timezoneOffset) * -1)
	endDateTime := startDateTime.Add(time.Hour * 24)

	semesterID, notFound, err := repository.QuerySemesterByFacultyTime(facultyID, startDateTime)
	if err != nil || notFound {
		return nil, err
	}

	courseIDInSemseterList, notFound, err := repository.QueryListCourseIDBySemester(semesterID)
	if err != nil || notFound {
		return nil, err
	}

	courseIDList, notFound, err := repository.QueryEnrollmentByListCourse(userID, courseIDInSemseterList)
	if err != nil || notFound {
		return nil, err
	}

	scheduleList, notFound, err := repository.QueryListScheduleByListCourseTime(courseIDList, startDateTime, endDateTime)
	if err != nil || notFound {
		return nil, err
	}

	var scheduleIDList []uint
	for i := 0; i < len(scheduleList); i++ {
		scheduleIDList = append(scheduleIDList, scheduleList[i].ID)
	}

	attendList, _, err := repository.QueryListAttendanceByUserSchedule(userID, scheduleIDList)
	if err != nil {
		return nil, err
	}

	var checkinTime *time.Time
	var checkinStatus string
	resultList := make([]dto.CheckInHistoryListElement, 0)
	for i := 0; i < len(scheduleList); i++ {

		checkinTime = nil
		checkinStatus = ""
		for j := 0; j < len(attendList); j++ {
			if scheduleList[i].ID == attendList[j].ScheduleID {
				checkinTime = &attendList[j].CheckInTime
				checkinStatus = attendList[j].CheckInStatus
			}
		}

		resultList = append(resultList, dto.CheckInHistoryListElement{
			Course:        scheduleList[i].Course.CourseID + " - " + scheduleList[i].Course.Name,
			StartTime:     scheduleList[i].StartTime,
			EndTime:       scheduleList[i].EndTime,
			Room:          scheduleList[i].Room.RoomID,
			CheckinTime:   checkinTime,
			CheckinStatus: checkinStatus,
		})
	}

	return resultList, nil
}
