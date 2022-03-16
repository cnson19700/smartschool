package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"
)

func GetUsers(ctx *context.Context) table.Table {
	tableUsers := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableUsers.GetInfo()

	info.AddField("ID", "id", db.Int).FieldHide()
	info.AddField("Email", "email", db.Varchar)

	info.AddButton("Import users", icon.FileExcelO, action.PopUp("/user", "Import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
				<div>
					<form id="form-import-excel" method="POST" action="/user" enctype="multipart/form-data">
						<input type="file" name="excel-file" id="file" accept=".xlsx" />
						<center>
							<input type="submit" value="Đăng tải"/>
						<center>
					</form>
				</div>`

			return true, "", data
		}))

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	return tableUsers
}
