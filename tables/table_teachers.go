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

func GetTeachers(ctx *context.Context) (tableTeachers table.Table) {
	tableTeachers = table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableTeachers.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Teacher", "teacher_name", db.Varchar)
	info.AddField("Teacher ID", "teacher_id", db.Varchar)

	info.HideNewButton()
	info.HideDeleteButton()
	info.HideEditButton()
	info.HideQueryInfo()
	info.AddCSS("span>.btn-group{display: none;}")
	info.SetTable("Teachers").SetTitle("Teachers").SetDescription("Teachers")

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetAllTeachers(param)
	})

	detail := tableTeachers.GetDetail()
	detail.AddField("Teacher", "teacher_name", db.Varchar)
	detail.AddField("Teacher ID", "teacher_id", db.Varchar)
	detail.AddField("Teaching Courses", "id", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		id, _ := value.Row["id"].(int)
		return template.
			Default().
			Link().
			SetURL("/admin/info/teacher_courses?__teacher_id=" + fmt.Sprint(id)).
			SetContent(template.HTML("Teaching Course(s)")).
			OpenInNewTab().
			SetTabTitle(template.HTML("Teaching Course(s)")).
			GetContent()
	})

	detail.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetTeacherData(param.GetFieldValue(parameter.PrimaryKey))
	})

	return
}

func GetAllTeachers(param parameter.Parameters) ([]map[string]interface{}, int) {
	sort := "desc"
	if len(param.SortType) > 0 {
		sort = param.SortType
	}
	pageSize := 5
	if len(param.PageSize) > 0 {
		pageSize = param.PageSizeInt
	}
	query := `
	select u.id, t.teacher_id, concat(u.first_name, ' ', u.last_name) as teacher_name
	from users u, teachers t
	where u.id = t.id
				ORDER BY u.id ` + sort

	var teacherResults []teacherResult
	database.DbInstance.Raw(query).Scan(&teacherResults)

	tableResults := make([]map[string]interface{}, len(teacherResults))
	for i, currentResult := range teacherResults {
		tempResult := make(map[string]interface{})

		tempResult["id"] = currentResult.ID
		tempResult["teacher_name"] = currentResult.TeacherName
		tempResult["teacher_id"] = currentResult.TeacherID

		tableResults[i] = tempResult
	}
	return tableResults, pageSize
}

func GetTeacherData(param string) ([]map[string]interface{}, int) {
	query := `
	select u.id, t.teacher_id, concat(u.first_name, ' ', u.last_name) as teacher_name
	from users u, teachers t, teacher_courses tc
	where u.id = ` + param + ` and u.role_id = 3 and u.id = t.id and t.id = tc.teacher_id
	order by u.id`
	var currentResult courseResult
	database.DbInstance.Raw(query).Scan(&currentResult)
	tableResult := make([]map[string]interface{}, 1)
	tempResult := make(map[string]interface{})

	tempResult["id"] = currentResult.ID
	tempResult["teacher_name"] = currentResult.TeacherName
	tempResult["teacher_id"] = currentResult.TeacherID
	tableResult[0] = tempResult

	return tableResult, 1
}

type teacherResult struct {
	ID          int
	TeacherName string
	TeacherID   string
}
