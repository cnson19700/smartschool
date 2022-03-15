package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"
)

func GetCourses(ctx *context.Context) table.Table {
	tableCourses := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableCourses.GetInfo()

	info.AddField("ID", "id", db.Int).FieldHide()
	info.AddField("Course", "course_id", db.Varchar)
	info.AddField("Name", "name", db.Varchar)

	info.AddButton("Import courses", icon.FileExcelO, action.PopUp("/course", "Import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
				<div>
					<form id="form-import-excel" method="POST" action="/course" enctype="multipart/form-data">
						<input type="file" name="excel-file" id="file" accept=".xlsx" />
						<center>
							<input type="submit" value="Đăng tải"/>
						<center>
					</form>
				</div>`

			return true, "", data
		}))

	info.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	return tableCourses
}
