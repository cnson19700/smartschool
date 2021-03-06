package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
)

func QueryCourseByID(id string) (*entity.Course, bool, error) {
	var course entity.Course
	result := database.DbInstance.Select("id").Where("course_id = ?", id).Find(&course)

	return &course, result.RowsAffected == 0, result.Error
}

func QueryAllCourses() (*[]entity.Course, error) {
	var course []entity.Course
	err := database.DbInstance.Find(&course).Error
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func QueryCourseBasicInfoByID(id uint) (*dto.CourseReportPartElement, bool, error) {
	var course dto.CourseReportPartElement
	result := database.DbInstance.Table("courses").Where("id = ?", id).Find(&course)

	return &course, result.RowsAffected == 0, result.Error
}

func QueryListCourseBySemester(sem_id uint) ([]dto.CourseReportPartElement, bool, error) {
	var queryList []dto.CourseReportPartElement
	result := database.DbInstance.Table("courses").Where("semester_id = ?", sem_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListCourseIDBySemester(sem_id uint) ([]uint, bool, error) {
	var queryList []uint
	result := database.DbInstance.Table("courses").Select("id").Where("semester_id = ? AND deleted_at IS NULL", sem_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListCourseBasicInfoByID(list_id []uint) ([]dto.CourseReportPartElement, bool, error) {
	var queryList []dto.CourseReportPartElement
	result := database.DbInstance.Table("courses").Where("id IN ? AND deleted_at IS NULL", list_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryCourseByTeacherID(teacher_id string) ([]*entity.CourseByTeacher, error) {
	courses := []*entity.Course{}
	courses_by_teacher := []*entity.CourseByTeacher{}

	err := database.DbInstance.Where("teacher_id = ?", teacher_id).Find(&courses).Error

	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		course := entity.CourseByTeacher{
			ID:              course.ID,
			CourseID:        course.CourseID,
			TeacherID:       course.TeacherID,
			TeacherRole:     course.TeacherRole,
			Name:            course.Name,
			SemesterID:      course.SemesterID,
			NumberOfStudent: course.NumberOfStudent,
		}
		courses_by_teacher = append(courses_by_teacher, &course)
	}

	return courses_by_teacher, nil
}

func QueryTeacherIDByCourseID(course_id uint) (uint, error) {
	var teacher_id uint
	result := database.DbInstance.Table("courses").Where("id = ?", course_id).Find(&teacher_id)

	return teacher_id, result.Error
}

func DeleteCourseByListCourseID(list_course_id []uint) error {
	var courses []entity.Course
	err := database.DbInstance.Delete(&courses, list_course_id).Error

	return err
}

func QueryCourseByCourseIdAndClass(courseId string, class string) (*entity.Course, bool, error) {
	var course entity.Course

	result := database.DbInstance.Where("course_id = ? and class = ?", courseId, class).
		Limit(1).Find(&course)

	return &course, result.RowsAffected == 0, result.Error
}

func QueryCourseIndexByCode(name string) int {
	ids := []int{}
	err := database.DbInstance.
		Table("courses").
		Select("courses.id").
		Where("courses.course_id = (?) and courses.deleted_at is null", name).
		Scan(&ids).Error
	if err != nil {
		return 0 // error
	}
	if len(ids) > 0 {
		return ids[0]
	}
	return 0 // not exist
}
