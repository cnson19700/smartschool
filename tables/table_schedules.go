package tables

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
	"gorm.io/gorm/clause"
)

func GetSchedules(ctx *context.Context) table.Table {
	tableSchedules := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableSchedules.GetInfo()
	info.HideNewButton()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Course Code", "course_name", db.Varchar)
	info.AddField("Room", "room_name", db.Varchar)
	info.AddField("Start Time", "start_time", db.Varchar)
	info.AddField("End Time", "end_time", db.Varchar)

	info.AddButton("Import Template", icon.FileExcelO, action.PopUp("/schedule", "Import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
				<div>
					<form id="form-import-excel" method="POST" action="/schedule" enctype="multipart/form-data">
						<input type="file" name="excel-file" id="file" accept=".xlsx" />
						<center>
							<input type="submit" value="Đăng tải"/>
						<center>
					</form>
				</div>`

			return true, "", data
		}))

	info.SetDeleteFn(func(ids []string) error {
		for _, id := range ids {
			if len(id) != 0 {
				var dbSchedule *entity.Schedule
				dbSchedule, _ = repository.QueryScheduleByID(id)

				if err := database.DbInstance.Delete(&dbSchedule).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetAllSchedulesData(param)
	})

	formList := tableSchedules.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Course Name", "course_id", db.Int, form.SelectSingle).FieldOptions(GetAllCoursesOption()).FieldDisplay(func(value types.FieldModel) interface{} {
		semester, _ := value.Row["course_name"].(string)
		return []string{semester}
	})
	formList.AddField("Room Name", "room_id", db.Int, form.SelectSingle).FieldOptions(GetAllRoomsData()).FieldDisplay(func(value types.FieldModel) interface{} {
		room, _ := value.Row["room_name"].(string)
		return []string{room}
	})
	formList.AddField("Start Time", "start_time", db.Timestamp, form.Datetime)
	formList.AddField("End Time", "end_time", db.Timestamp, form.Datetime)
	formList.EnableAjax("Success", "Fail")
	formList.SetTable("schedules").SetTitle("Schedules").SetDescription("Schedules")

	formList.SetUpdateFn(func(values form2.Values) error {
		if values.IsEmpty("course_id") {
			return errors.New("Course cannot be empty")
		}
		if values.IsEmpty("room_id") {
			return errors.New("Room cannot be empty")
		}
		if values.IsEmpty("start_time") {
			return errors.New("Start Time cannot be empty")
		}
		if values.IsEmpty("end_time") {
			return errors.New("End Time cannot be empty")
		}
		id, _ := strconv.Atoi(values.Get("id"))

		updated := database.DbInstance.Model(&entity.Schedule{}).Where("id = ? ", id).Clauses(clause.Returning{}).Updates(map[string]interface{}{
			"room_id":    values.Get("room_id"),
			"course_id":  values.Get("course_id"),
			"start_time": values.Get("start_time") + "+7:00",
			"end_time":   values.Get("end_time") + "+7:00",
		}).Error
		if updated != nil {
			return updated
		}
		return nil
	})

	detail := tableSchedules.GetDetail()
	detail.AddField("Course Name", "course_name", db.Varchar)
	detail.AddField("Room Name", "room_name", db.Varchar)
	detail.AddField("Start Time", "start_time", db.Timestamp)
	detail.AddField("End Time", "end_time", db.Timestamp)

	detail.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetScheduleData(param.GetFieldValue(parameter.PrimaryKey))
	})

	return tableSchedules
}

func GetAllSchedulesData(param parameter.Parameters) ([]map[string]interface{}, int) {
	query := `
	select sc.id, r.name as room_name, c.name as course_name, sc.start_time, sc.end_time
		from schedules sc
		left join rooms r
		on r.id = sc.room_id
		left join courses c
		on sc.course_id = c.id
		where sc.deleted_at is null
		order by sc.id`

	var scheduleResults []scheduleResult
	database.DbInstance.Raw(query).Scan(&scheduleResults)
	tableResults := make([]map[string]interface{}, len(scheduleResults))
	for i, current_schedule := range scheduleResults {
		tempResult := make(map[string]interface{})

		tempResult["id"] = current_schedule.ID
		tempResult["course_name"] = current_schedule.CourseName
		tempResult["room_name"] = current_schedule.RoomName
		tempResult["start_time"] = current_schedule.StartTime.Format("15:04:05")
		tempResult["end_time"] = current_schedule.EndTime.Format("15:04:05")

		tableResults[i] = tempResult
	}
	return tableResults, 10
}

func GetScheduleData(param string) ([]map[string]interface{}, int) {
	query := `
	select sc.id, r.name as room_name, c.name as course_name, sc.start_time, sc.end_time
		from schedules sc
		left join rooms r
		on r.id = sc.room_id
		left join courses c
		on sc.course_id = c.id
		where sc.id = ` + param + `
		order by sc.id`

	var current_schedule scheduleResult
	database.DbInstance.Raw(query).Scan(&current_schedule)
	tableResults := make([]map[string]interface{}, 1)
	tempResult := make(map[string]interface{})

	tempResult["id"] = current_schedule.ID
	tempResult["course_name"] = current_schedule.CourseName
	tempResult["room_name"] = current_schedule.RoomName
	tempResult["start_time"] = current_schedule.StartTime.Format("15:04:05")
	tempResult["end_time"] = current_schedule.EndTime.Format("15:04:05")

	tableResults[0] = tempResult
	return tableResults, 1
}

func GetAllRoomsData() []types.FieldOption {
	room_options := []types.FieldOption{}

	room_results := []RoomOptionResult{}

	query := `select s.id, s.name
				from rooms s`

	database.DbInstance.Raw(query).Scan(&room_results)
	for _, r := range room_results {
		tmp := types.FieldOption{
			Text:  r.Name,
			Value: fmt.Sprint(r.ID),
		}
		room_options = append(room_options, tmp)
	}
	return room_options
}

func GetAllCoursesOption() []types.FieldOption {
	course_options := []types.FieldOption{}

	course_results := []CourseOptionResult{}

	query := `select s.id, s.name as name
 				from courses s`

	database.DbInstance.Raw(query).Scan(&course_results)
	for _, r := range course_results {
		tmp := types.FieldOption{
			Text:  r.Name,
			Value: fmt.Sprint(r.ID),
		}
		course_options = append(course_options, tmp)
	}
	return course_options
}

type RoomOptionResult struct {
	ID   int
	Name string
}
type CourseOptionResult struct {
	ID   int
	Name string
}

type scheduleResult struct {
	ID         int
	RoomName   string
	CourseName string
	StartTime  time.Time
	EndTime    time.Time
}
