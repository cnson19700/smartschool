package service

import (
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
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

func GetCheckInHistoryBySID(sid string, status string) (*entity.Student, []dto.CheckInHistoryElement) {

	// student, err := repo.QueryStudentBySID(sid)
	// if err != nil {
	// 	return nil, nil
	// }

	// listHistory := repo.QueryCheckinHistoryWithSIdAndStatus(student.ID, status)

	// var historyElements = make([]dto.CheckInHistoryElement, 0)
	// for i := 0; i < len(listHistory); i++ {
	// 	historyElements = append(historyElements, dto.CheckInHistoryElement{
	// 		CourseName:    listHistory[i].Scheduler.Course.CourseID + " - " + listHistory[i].Scheduler.Course.Name,
	// 		CheckinTime:   listHistory[i].CheckInTime,
	// 		CheckinStatus: listHistory[i].CheckInStatus})
	// }

	// return student, historyElements
	return nil, nil
}

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
		Name:        user.Username,
		Class:       student.Batch,
		Email:       user.Email,
		Gender:      genderStudent,
		PhoneNumber: user.PhoneNumber,
	}
	return &StudentProfile, nil
}
