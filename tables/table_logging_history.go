package tables

import (
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/smartschool/database"
	"github.com/smartschool/repository"
)

func GetLoggingHistories(ctx *context.Context) (tableLoggings table.Table) {
	tableLoggings = table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableLoggings.GetInfo().HideRowSelector()
	info.HideDeleteButton()
	info.HideNewButton()
	info.HideEditButton()

	info.AddField("ID", "id", db.Int)
	info.AddField("Device ID", "device_id", db.Varchar)
	info.AddField("Time", "created_at", db.Timestamp)
	info.AddField("User ID", "user_id", db.Text).FieldWidth(300)
	info.AddColumnButtons("User Detail", types.GetColumnButton("View User Info", icon.Info,
		action.PopUp("/user_id_detail", "Details", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			device_signal_logs, _, _ := repository.QueryUserIDByDeviceSignalID(ctx.FormValue("id"))
			student, err := repository.QueryStudentByID(fmt.Sprint(device_signal_logs.CardId))
			room_id, _, _ := repository.QueryRoomByDeviceID(device_signal_logs.CompanyTokenKey)
			room, _, _:= repository.QueryRoomInfo(room_id)
			if (err != nil){
				return true, "Invalid Record", `<h2 id='room_title'>Room: `+room.RoomID+`</h2><h3 id='status_title'>`+ device_signal_logs.Status + `</h3> <style>
					h2#room_title, h3#status_title {text-align: center;}
				</style>`
			} else{
			return true, "ok", "<h2 id='student_title'>"+ student.StudentID +"</h2> <h4>Room: "+ room.RoomID +"</h4> <h4>Status: "+ device_signal_logs.Status+`</h4>`
		}
		
		})))

	info.HideEditButton()
	info.HideDeleteButton()
	info.HideDetailButton()

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetLogsData()
	})

	return
}

func GetLogsData()([]map[string]interface{}, int){
	query := `
	select d.id, d.card_id as user_id, d.created_at, d.status, d.company_token_key as device_id
	from device_signal_logs d
	inner join devices
	on device_id = company_token_key
	`
	var currentResult []logResult
	database.DbInstance.Raw(query).Scan(&currentResult)
	tableResult := make([]map[string]interface{}, len(currentResult))
	for i, current := range currentResult{
		tempResult := make(map[string]interface{})

		tempResult["id"] = current.ID
		tempResult["device_id"] = current.DeviceID
		tempResult["created_at"] = current.CreatedAt.Format("2006-01-02 15:04:05")
		tempResult["status"] = current.Status
		tempResult["user_id"] = current.UserID

		tableResult[i] = tempResult
	}

	return tableResult, 10
}

type logResult struct{
	ID int
	DeviceID string
	UserID string
	Status string
	CreatedAt time.Time
}