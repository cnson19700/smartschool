package tables

import (
	"fmt"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

func GetAbsenceStudents(ctx *context.Context) table.Table {
	tableAbsenceStudents := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))
	info := tableAbsenceStudents.GetInfo().HideFilterArea()

	info.HideNewButton()
	info.HideDetailButton()
	info.HideEditButton()
	info.HideDeleteButton()
	info.HideQueryInfo()
	info.HideExportButton()
	info.HideFilterButton()
	info.HideQueryInfo()

	info.AddField("ID", "id", db.Int)
	info.AddField("Student Name", "student_name", db.Varchar)
	info.AddField("Student ID", "user_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar)

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetAbsencesData(param)
	})
	return tableAbsenceStudents
}

func GetAbsencesData(params parameter.Parameters) ([]map[string]interface{}, int) {
	course, _, _ := repository.QueryCourseByID(params.GetFieldValue("course_id"))
	created_query := params.GetFieldValue("created_at_start__goadmin")
	course_id := fmt.Sprint(course.ID)
	schedules := params.GetFieldValue("schedule_ids")
	var query string
	studentIds := []uint{}
	absenceStudents := []*entity.User{}

	// date_constaint := `and a.created_at BETWEEN NOW()::DATE-EXTRACT(DOW FROM NOW())::INTEGER-7
	// AND NOW()::DATE-EXTRACT(DOW from NOW())::INTEGER))`

	date_constaint := `and a.created_at >= current_date at time zone 'UTC' - interval '6 days')`
	schedules_ids := strings.Split(schedules, ",")

	if len(schedules_ids) > 1 {
		query = `select s.student_id
	from student_course_enrollments s
	where s.course_id = ` + course_id + ` and s.id not in 
	(select st.id
	from student_course_enrollments st
	left join attendances a
	on st.student_id = a.user_id
	where st.course_id = ` + course_id + ` and a.schedule_id in (` + schedules + `) `
	} else {
		query = `select s.student_id
		from student_course_enrollments s
		where s.course_id = ` + course_id + ` and s.id not in 
		(select st.id
		from student_course_enrollments st
		left join attendances a
		on st.student_id = a.user_id
		where st.course_id = ` + course_id + ` and a.schedule_id = ` + schedules
	}

	if len(created_query) > 0 {
		query += ` and a.created_at >= '` + created_query + "')"
	} else {
		query += date_constaint
	}
	database.DbInstance.Raw(query).Scan(&studentIds)
	query_students := database.DbInstance.Model(&entity.User{})
	if len(studentIds) > 1 {
		query_students.Where("id in ? ", studentIds)
	} else {
		query_students.Where("id in = ", studentIds)
	}

	query_students.Order("Id").Find(&absenceStudents)

	absence_results := make([]map[string]interface{}, len(absenceStudents))

	for i, currentResult := range absenceStudents {
		tempResult := make(map[string]interface{})

		tempResult["id"] = currentResult.ID
		tempResult["student_name"] = currentResult.LastName + " " + currentResult.FirstName
		tempResult["user_name"] = currentResult.Username
		tempResult["email"] = currentResult.Email

		absence_results[i] = tempResult
	}

	return absence_results, 10
}
