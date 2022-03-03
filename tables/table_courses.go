package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCourses(ctx *context.Context) table.Table {
	tableCourses := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableCourses.GetInfo()

	info.AddField("ID", "id", db.Int).FieldHide()
	info.AddField("Course", "course_id", db.Varchar)
	info.AddField("Name", "name", db.Varchar)

	info.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	return tableCourses
}
