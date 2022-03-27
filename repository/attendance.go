package repository

import (
	"fmt"
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByTeacherCourse(teacherID string, courseId string) ([]*entity.Attendance, error) {
	attendance := []*entity.Attendance{}
	err := database.DbInstance.Where("teacher_id = ? AND course_id =?", teacherID, courseId).Find(&attendance).Error
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func QueryAttendanceByCourseID(courseId string) ([]*entity.Attendance, error) {
	attendance := []*entity.Attendance{}
	err := database.DbInstance.Where("course_id =?", courseId).Find(&attendance).Error
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func QueryAttendanceByStudentSchedule(student_id string, schedule_id uint) (*entity.Attendance, error) {
	var checkAttend entity.Attendance
	result := database.DbInstance.Select("id").Where("user_id = ? AND schedule_id = ?", student_id, schedule_id).Find(&checkAttend)

	return &checkAttend, result.Error
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

func SearchAttendance(pagnitor *entity.Paginator, filter *entity.AttendanceFilter,
	orders []string) ([]*entity.Attendance, error) {
	query := database.DbInstance.Model(&entity.Attendance{})

	//Order
	for _, order := range orders {
		query.Order(order)
	}

	fmt.Println(filter)

	if filter.Keyword != "" { //search   checkin_time filter select?
		query.Where("title LIKE ?", "%"+filter.Keyword+"%") //user_name, schedule_id, user_id, checkin_status
	}
}
