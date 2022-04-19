package repository

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByTeacherCourse(teacherID string, courseId string) ([]*entity.AttendanceResult, error) {
	attendances := []*entity.Attendance{}
	scheduleIDs := []uint{}
	attendance_results := []*entity.AttendanceResult{}

	database.DbInstance.Table("schedules").Select("id").Where("course_id = ?", courseId).Scan(&scheduleIDs)

	err := database.DbInstance.Where("teacher_id = ? AND schedule_id IN ?", teacherID, scheduleIDs).Find(&attendances).Error
	if err != nil {
		return nil, err
	}

	for _, attendance := range attendances {
		student, _ := QueryStudentByID(fmt.Sprint(attendance.UserID))
		user := QueryUserBySID(fmt.Sprint(attendance.UserID))
		attendance_result := &entity.AttendanceResult{
			ID:            attendance.ID,
			TeacherID:     attendance.TeacherID,
			StudentID:     student.StudentID,
			StudentName:   user.FirstName + " " + user.LastName,
			ScheduleID:    attendance.ScheduleID,
			CheckinStatus: attendance.CheckInStatus,
		}
		attendance_results = append(attendance_results, attendance_result)
	}

	return attendance_results, nil
}

// func QueryAttendanceByTeacherCourse(teacherID string, courseId string) ([]*entity.Attendance, error) {
// 	attendance := []*entity.Attendance{}
// 	err := database.DbInstance.Where("teacher_id = ? AND course_id =?", teacherID, courseId).Find(&attendance).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return attendance, nil
// }

func QueryAttendanceByCourseID(courseId string) ([]*entity.Attendance, error) {
	attendance := []*entity.Attendance{}
	err := database.DbInstance.Where("course_id =?", courseId).Find(&attendance).Error
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func QueryAttendanceByStudentSchedule(student_id string, schedule_id uint) (bool, error) {
	var checkAttend entity.Attendance
	result := database.DbInstance.Select("id").Where("user_id = ? AND schedule_id = ?", student_id, schedule_id).Find(&checkAttend)

	return result.RowsAffected == 0, result.Error
}

func CreateAttendance(attendance entity.Attendance) error {
	err := database.DbInstance.Create(&attendance).Error

	return err
}

func QueryListAttendanceByUserSchedule(user_id uint, schedule_id_list []uint) ([]entity.Attendance, bool, error) {
	var queryList []entity.Attendance
	result := database.DbInstance.Where("user_id = ? AND schedule_id IN ?", user_id, schedule_id_list).Find(&queryList)

	//attendanceList := append([]entity.Attendance{}, queryList...)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListAttendanceInDayByUser(user_id uint, start time.Time, end time.Time) ([]entity.Attendance, bool, error) {
	var queryList []entity.Attendance
	result := database.DbInstance.Where("user_id = ? AND (checkin_time BETWEEN ? AND ?)", user_id, start, end).Preload("Schedule.Room").Preload("Schedule.Course").Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func CountAttendanceOfSchedule(user_id uint, schedule_id_list []uint) (int64, error) {
	var c int64
	result := database.DbInstance.Table("attendances").Select("id").Where("user_id = ? AND schedule_id IN ?", user_id, schedule_id_list).Count(&c)

	return c, result.Error
}
  
func SearchAttendance(params url.Values) ([]*entity.AttendanceResult, error) {
	//params includes student_id, student_name, order[], checkin_status, checkin_day
	filter := entity.AttendanceFilter{
		StudentName:   strings.ToLower(params.Get("student_name")),
		StudentID:     params.Get("student_id"),
		CheckinStatus: strings.ToLower(params.Get("checkin_status")),
		CheckinDay:    params.Get("checkin_day"),
	}

	spew.Dump(filter)
	orders := strings.Split(params.Get("orders"), ",")
	query := database.DbInstance.Model(&entity.Attendance{})
	attendances := []*entity.Attendance{}
	attendance_results := []*entity.AttendanceResult{}

	if filter.StudentID != "" {
		student, _, _ := QueryStudentBySID(filter.StudentID)
		query.Where("user_id = ?", student.ID)
	}
	if filter.StudentName != "" {
		student_ids, _ := QueryStudentsByName(filter.StudentName) //return user_ids
		if len(student_ids) > 1 {
			query.Where("user_id IN ? ", student_ids)
		} else {
			fmt.Printf("%v\n", len(student_ids))
			query.Where("user_id = ?", student_ids[0])
		}
	}
	if filter.CheckinStatus != "" {
		query.Where("LOWER(checkin_status) LIKE ? ", "%"+filter.CheckinStatus+"%")
	}

	// if filter.CheckinDay != "" {

	// }

	for _, order := range orders {
		query.Order(order)
	}
	err := query.Find(&attendances).Error

	if err != nil {
		return nil, err
	}

	for _, attendance := range attendances {
		student, _ := QueryStudentByID(fmt.Sprint(attendance.UserID))
		user := QueryUserBySID(fmt.Sprint(attendance.UserID))
		attendance_result := &entity.AttendanceResult{
			ID:            attendance.ID,
			TeacherID:     attendance.TeacherID,
			StudentID:     student.StudentID,
			StudentName:   user.FirstName + " " + user.LastName,
			ScheduleID:    attendance.ScheduleID,
			CheckinStatus: attendance.CheckInStatus,
		}

		attendance_results = append(attendance_results, attendance_result)
	}

	return attendance_results, nil
}
