package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
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

func ExistAttendanceByUserSchedule(student_id uint, schedule_id uint) (bool, error) {
	var checkAttendID uint
	result := database.DbInstance.Table("attendances").Select("id").Where("user_id = ? AND schedule_id = ? AND deleted_at IS NULL", student_id, schedule_id).Find(&checkAttendID)

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
	result := database.DbInstance.Table("attendances").Select("id").Where("user_id = ? AND schedule_id IN ? AND deleted_at IS NULL", user_id, schedule_id_list).Count(&c)

	return c, result.Error
}

func SearchAttendance(params parameter.Parameters) ([]*entity.Attendance, error) {
	//params includes student_id, student_name, order[], checkin_status, checkin_day
	filter := entity.AttendanceFilter{
		StudentName:     strings.ToLower(params.GetFieldValue("student_name")),
		StudentID:       params.GetFieldValue("student_id"),
		CheckinStatus:   strings.ToLower(params.GetFieldValue("checkin_status")),
		CheckinDayStart: params.GetFieldValue("created_at_start__goadmin"),
		CheckinDayTo:    params.GetFieldValue("created_at_end__goadmin"),
	}

	order := params.SortType
	query := database.DbInstance.Model(&entity.Attendance{})
	attendances := []*entity.Attendance{}
	scheduleIDs := []uint{}
	teacher_id, course_id := params.GetFieldValue("teacher_id"), params.GetFieldValue("course_id")
	if len(teacher_id) > 0 && len(course_id) > 0 {
		database.DbInstance.Table("schedules").Select("id").Where("course_id = ?", course_id).Scan(&scheduleIDs)
		if len(scheduleIDs) > 1 {
			query.Where("teacher_id = ? AND schedule_id IN ? ", teacher_id, scheduleIDs)
		} else {
			query.Where("teacher_id = ? AND schedule_id = ? ", teacher_id, scheduleIDs)
		}
	}

	if filter.StudentID != "" {
		student_ids := []uint{}
		database.DbInstance.Table("students").Select("students.id").Where("student_id LIKE ? ", "%"+filter.StudentID+"%").Scan(&student_ids)
		query.Where("user_id IN ? ", student_ids)
	}
	if filter.StudentName != "" {
		student_ids, _ := QueryStudentsByName(filter.StudentName) //return user_ids
		if len(student_ids) > 1 {
			query.Where("user_id IN ? ", student_ids)
		} else {
			query.Where("user_id = ? ", student_ids[0])
		}
	}
	if filter.CheckinStatus != "" {
		query.Where("LOWER(checkin_status) LIKE ? ", "%"+filter.CheckinStatus+"%")
	}

	if filter.CheckinDayStart != "" {
		if filter.CheckinDayTo != "" {
			query.Where("created_at BETWEEN ? AND ? ", filter.CheckinDayStart, filter.CheckinDayTo)
		} else {
			query.Where("created_at BETWEEN ? AND ?", filter.CheckinDayStart, time.Now())
		}
	}

	if params.GetFieldValue("batch") != "" {

	}

	query.Order("id " + order)
	err := query.Find(&attendances).Error

	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func AttendanceByTeacherCourse(params parameter.Parameters) ([]*entity.Attendance, error) {
	filter := entity.AttendanceFilter{
		StudentName:     strings.ToLower(params.GetFieldValue("student_name")),
		StudentID:       params.GetFieldValue("student_id"),
		CheckinStatus:   strings.ToLower(params.GetFieldValue("checkin_status")),
		CheckinDayStart: params.GetFieldValue("created_at_start__goadmin"),
		CheckinDayTo:    params.GetFieldValue("created_at_end__goadmin"),
	}

	// order := params.SortType
	query := database.DbInstance.Model(&entity.Attendance{})
	attendances := []*entity.Attendance{}
	scheduleIDs := []uint{}
	teacher_id, course_id := params.GetFieldValue("teacher_id"), params.GetFieldValue("course_id")
	if len(teacher_id) > 0 && len(course_id) > 0 {
		// database.DbInstance.Table("schedules").Select("id").Where("course_id = ?", course_id).Scan(&scheduleIDs)
		// if len(scheduleIDs) > 1 {
		// 	query.Where("teacher_id = ? AND schedule_id IN ? ", teacher_id, scheduleIDs)
		// } else {
		// 	query.Where("teacher_id = ? AND schedule_id = ? ", teacher_id, scheduleIDs)
		// }
		query_schedules := `select distinct schedules.id
		from schedules
		inner join 
		(select distinct courses.*
		from courses
		inner join teacher_courses on teacher_courses.course_id = courses.id
		where teacher_courses.teacher_id = ` + teacher_id + ` and 
		courses.course_id = '` + course_id + `') c
			on c.id = schedules.course_id`

		database.DbInstance.Raw(query_schedules).Scan(&scheduleIDs)
		if len(scheduleIDs) > 1 {
			query.Where("schedule_id IN ? ", scheduleIDs)
		} else {
			query.Where("schedule_id = ? ", scheduleIDs)
		}
	} else if len(course_id) > 0 {
		query_schedules := `select distinct schedules.id
		from schedules
		inner join 
		(select distinct courses.*
		from courses
		inner join teacher_courses on teacher_courses.course_id = courses.id
		where courses.course_id = '` + course_id + `') c
			on c.id = schedules.course_id`

		database.DbInstance.Raw(query_schedules).Scan(&scheduleIDs)
		if len(scheduleIDs) > 1 {
			query.Where("schedule_id IN ? ", scheduleIDs)
		} else {
			query.Where("schedule_id = ? ", scheduleIDs)
		}
	} else {
		query.Where("user_id = null")
	}

	if filter.StudentID != "" {
		student_ids := []uint{}
		database.DbInstance.Table("students").Select("students.id").Where("student_id LIKE ? ", "%"+filter.StudentID+"%").Scan(&student_ids)
		query.Where("user_id IN ? ", student_ids)
	}
	if filter.StudentName != "" {
		student_ids, _ := QueryStudentsByName(filter.StudentName) //return user_ids
		if len(student_ids) > 1 {
			query.Where("user_id IN ? ", student_ids)
		} else if len(student_ids) == 1 {
			query.Where("user_id = ? ", student_ids[0])
		} else {
			query.Where("user_id = null")
		}
	}
	if filter.CheckinStatus != "" {
		query.Where("LOWER(checkin_status) LIKE ? ", "%"+filter.CheckinStatus+"%")
	}
	if filter.CheckinDayStart != "" {
		if filter.CheckinDayTo != "" {
			query.Where("created_at BETWEEN ? AND ? ", filter.CheckinDayStart, filter.CheckinDayTo)
		} else {
			query.Where("created_at BETWEEN ? AND ?", filter.CheckinDayStart, time.Now())
		}
	}

	err := query.Order("created_at DESC").Find(&attendances).Error

	if err != nil {
		return nil, err
	}
	return attendances, nil
}
