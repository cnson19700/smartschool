package tables

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

func GetAttendances(ctx *context.Context) table.Table {

	tableAttendaces := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableAttendaces.GetInfo()
	info.AddField("ID", "id", db.Int).HideEditButton().FieldSortable()
	info.AddField("Student ID", "student_id", db.Varchar)
	info.AddField("Student Name", "student_name", db.Varchar)
	info.AddField("Checkin Status", "checkin_status", db.Varchar)
	info.AddField("Created At", "created_at", db.Timestamp)
	info.HideNewButton()
	info.HideDetailButton()
	info.HideDeleteButton()

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		teacher_id, course_id := "", ""
		if len(param.GetFieldValue("teacher_id")) > 0 && len(param.GetFieldValue("course_id")) > 0 {
			teacher_id, course_id = param.GetFieldValue("teacher_id"), param.GetFieldValue("course_id")
		}
		return GetAllAttendancesData(teacher_id, course_id)
	})

	info.AddButton("Import attendances", icon.FileExcelO, action.PopUp("/attendance", "Import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
				<div>
					<form id="form-import-excel" method="POST" action="/attendance" enctype="multipart/form-data">
						<input type="file" name="excel-file" id="file" accept=".xlsx" />
						<center>
							<input type="submit" value="Đăng tải"/>
						<center>
					</form>
				</div>`

			return true, "", data
		}))
	info.SetTable("attendances").SetTitle("Attendances").SetDescription("Attendances")
	return tableAttendaces
}

func GetAllAttendancesData(teacher_id string, course_id string) ([]map[string]interface{}, int) {
	attendances := []*entity.Attendance{}
	scheduleIDs := []uint{}
	if len(teacher_id) > 0 && len(course_id) > 0 {
		database.DbInstance.Table("schedules").Select("id").Where("course_id = ?", course_id).Scan(&scheduleIDs)
		database.DbInstance.Where("teacher_id = ? AND schedule_id IN ?", teacher_id, scheduleIDs).Find(&attendances)
	} else {
		database.DbInstance.Find(&attendances)
	}

	attendance_results := make([]map[string]interface{}, len(attendances))

	for i, attendance := range attendances {
		student, _ := repository.QueryStudentByID(fmt.Sprint(attendance.UserID))
		user := repository.QueryUserBySID(fmt.Sprint(attendance.UserID))
		attendance_result := make(map[string]interface{})

		attendance_result["id"] = i
		attendance_result["teacher_id"] = attendance.TeacherID
		attendance_result["student_id"] = student.StudentID
		attendance_result["student_name"] = user.FirstName + " " + user.LastName
		attendance_result["schedule_id"] = attendance.ScheduleID
		attendance_result["checkin_status"] = attendance.CheckInStatus

		attendance_results[i] = attendance_result
	}

	return attendance_results, 10
}
