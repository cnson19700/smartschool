package tables

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/smartschool/database"
)

func GetTeacherCourses(ctx *context.Context) (tableTeacherCourses table.Table) {
	tableTeacherCourses = table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableTeacherCourses.GetInfo()
	info.AddField("CourseID", "course_id", db.Varchar)
	info.AddField("Course Name", "id", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		course_id, _ := value.Row["course_id"].(string)
		return template.
			Default().
			Link().
			SetURL("/admin/info/courses/detail?__goadmin_detail_pk=" + fmt.Sprint(course_id)).
			SetContent(template.HTML(value.Row["course_name"].(string))).
			OpenInNewTab().
			SetTabTitle(template.HTML(value.Row["course_name"].(string))).
			GetContent()
	})
	info.AddField("Class", "class", db.Varchar)
	info.AddField("Role in course", "teacher_role", db.Varchar)
	info.AddField("Semester", "semester_name", db.Varchar)
	info.AddField("Attendance", "attendance", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		teacher_id, _ := value.Row["teacher_id"].(string)
		course_code, _ := value.Row["course_id"].(string)
		class, _ := value.Row["class"].(string)
		return template.
			Default().
			Link().
			SetURL("/admin/info/attendances?teacher_id=" + teacher_id + "&course_id=" + course_code + "&class=" + class).
			SetContent("Attendances").
			OpenInNewTab().
			GetContent()
	})
	info.HideEditButton()
	info.HideDeleteButton()

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetTeacherCoursesData(param.GetFieldValue("__teacher_id"))
	})

	return
}

func GetTeacherCoursesData(param string) ([]map[string]interface{}, int) {
	query := `
	select c.id, c.course_id, c.name as course_name, c.class, s.title as semester_name, s.id as semester_id, tc.teacher_role, tc.teacher_id
	from courses c, teacher_courses tc, semesters s
	where tc.teacher_id = ` + param + ` and tc.course_id = c.id and s.id = c.semester_id
	order by c.id`
	var currentResult []courseResult
	database.DbInstance.Raw(query).Scan(&currentResult)
	tableResult := make([]map[string]interface{}, len(currentResult))
	for i, currentResult := range currentResult {
		tempResult := make(map[string]interface{})

		tempResult["id"] = currentResult.ID
		tempResult["course_id"] = currentResult.CourseID
		tempResult["course_name"] = currentResult.CourseName
		tempResult["class"] = currentResult.Class
		tempResult["semester_name"] = currentResult.SemesterName
		tempResult["semester_id"] = currentResult.SemesterID
		tempResult["teacher_role"] = currentResult.TeacherRole
		tempResult["teacher_id"] = currentResult.TeacherID
		tableResult[i] = tempResult
	}

	return tableResult, 1
}
