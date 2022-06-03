package tables

import (
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/davecgh/go-spew/spew"
	"github.com/smartschool/database"
)

func GetLoggingHistories(ctx *context.Context) (tableLoggings table.Table) {
	tableLoggings = table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableLoggings.GetInfo().HideRowSelector()
	info.AddField("ID", "id", db.Int)
	info.AddField("Device ID", "device_id", db.Varchar)
	info.AddField("Time", "created_at", db.Timestamp)
	info.AddField("User ID", "user_id", db.Text).FieldWidth(300)

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

		spew.Dump(tempResult)
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