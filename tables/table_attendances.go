package tables

import (
	"fmt"
	"strings"

	template2 "html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

var total_lates, total_intime, total_absences, total_student_inclass = 0, 0, 0, 0
var batch string
var in_time_schedules []*entity.Schedule

func GetAttendances(ctx *context.Context) table.Table {

	tableAttendaces := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	tmp_course, _, _ := repository.QueryCourseByID(ctx.FormValue("course_id"))
	student_id, _ := repository.QueryFirstStudentInCourseEnrollment(int(tmp_course.ID))
	student, _ := repository.QueryStudentByID(fmt.Sprint(student_id))
	batch = student.Batch

	info := tableAttendaces.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.Int).HideEditButton().FieldSortable()
	info.AddField("Course", "course_id", db.Varchar).FieldFilterable(types.FilterType{
		Options: []types.FieldOption{
			{
				Text:     ctx.FormValue("course_id"),
				Value:    ctx.FormValue("course_id"),
				Selected: true,
			},
		},
		FormType:    form.SelectSingle,
		NoIcon:      false,
		Placeholder: ctx.FormValue("course_id"),
	}).HideFilterButton().FieldHide()
	info.AddField("Teacher ID", "teacher_id", db.Varchar).FieldFilterable(types.FilterType{
		Options: []types.FieldOption{
			{Text: ctx.FormValue("teacher_id"),
				Value:    ctx.FormValue("teacher_id"),
				Selected: true,
			},
		},
		FormType:    form.SelectSingle,
		NoIcon:      false,
		Placeholder: ctx.FormValue("teacher_id"),
	}).FieldHide()
	info.AddField("Batch", "batch", db.Varchar).FieldHide()
	info.AddField("Student ID", "student_id", db.Varchar).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(GetAllStudentIDs())
	info.AddField("Student Name", "student_name", db.Varchar).FieldFilterable()
	info.AddField("Checkin Status", "checkin_status", db.Varchar).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "Late", Text: "Late"},
		{Value: "Attend", Text: "Attend"},
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		c, _ := value.Row["checkin_status"].(string)
		return c
	})
	info.AddField("Created At", "created_at", db.Timestamp).FieldFilterable(types.FilterType{FormType: form.DateRange, Placeholder: " ... "}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			c, _ := value.Row["created_at"].(string)
			var a, b int
			var d float64
			add_attr := "attend-stt"
			database.DbInstance.Raw("select extract(dow from date '" + strings.Split(c, " ")[0] + "')").Scan(&a)
			for _, i := range in_time_schedules {
				database.DbInstance.Raw("select extract(dow from date '" + i.StartTime.Format("2006-01-02") + "')").Scan(&b)
				if a == b {
					database.DbInstance.Raw("select extract(epoch FROM (TIMESTAMP '" + c + "' -  TIMESTAMP'" + i.StartTime.Format("2006-01-02 15:04:05") + "'))/60").Scan(&d)
					//checkin time - schedule time > 0 == late
					if int(d) > 0 && int(d) < 10 {
						add_attr = "range_five"
					} else if int(d) >= 10 {
						add_attr = "range_ten"
					}
				}
			}
			return "<span id=" + add_attr + ">" + c + "</span>"
		})
	info.HideNewButton()
	info.HideDetailButton()
	info.HideDeleteButton()
	info.HideQueryInfo()
	info.AddCSS(`
		#range_ten{color: #8B0000;} 
		#attend-stt{color: #228B22;} 
		#range_five{color: #FF8C00;} 
		.clearfix{display: none;}`)
	info.AddJS(`
		row = $('tr > td > span#range_five').parent();
		$(row).parent().attr("id", "range_five");

		row = $('tr > td > span#range_ten').parent();
		$(row).parent().attr("id", "range_ten");

		row = $('tr > td > span#attend-stt').parent();
		$(row).parent().attr("id", "attend-stt");
	`)

	info.AddCSS(".reset {visibility: hidden;}  span>.btn-group{display: none;}")

	info.SetTable("Overview").SetTitle("Overview").SetDescription("Overview").
		SetWrapper(func(content template2.HTML) template2.HTML {
			col1 := `<div style="margin-left:243px;">` + content + `</div>`

			table_sum := template.Default().Table().SetThead(types.Thead{
				{
					Head: "Status",
				},
				{
					Head: "Times",
				},
			}).SetInfoList([]map[string]types.InfoItem{
				{"Status": types.InfoItem{Content: "Lates"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_lates) + "/" + fmt.Sprint(total_student_inclass))}},
				{"Status": types.InfoItem{Content: "In times"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_intime) + "/" + fmt.Sprint(total_student_inclass))}},
				// {"Title": types.InfoItem{Content: "Absences"},
				// 	"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_absences))}},
			},
			).SetMinWidth("100px").GetContent()
			col2 := `<div style="position: absolute;width:230px;">` + template.Default().Box().SetHeader("Overview").
				WithHeadBorder().SetBody(table_sum).GetContent() + `</div>`
			return `<div style="width:100%;">` + col2 + col1 + `</div>`
		})
	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetAllAttendancesData(param) //base on teacher_course
	})
	created_at_query := ctx.FormValue("created_at_start__goadmin")
	absence_url :=
		"/admin/info/absence_students?course_id=" + ctx.FormValue("course_id") +
			"&schedule_ids=" + getScheduleIds(in_time_schedules) + "&created_at_start__goadmin=" + created_at_query

	info.AddButton("Absences Students", icon.Tv, action.Jump(absence_url))
	info.SetTable("attendances").SetTitle("Attendances " + ctx.FormValue("course_id") + " - " + batch).SetDescription("Attendances")
	return tableAttendaces
}

func GetAllAttendancesData(param parameter.Parameters) ([]map[string]interface{}, int) {
	total_absences, total_intime, total_lates = 0, 0, 0
	attendances := []*entity.Attendance{}
	attendances, _, in_time_schedules = repository.AttendanceByTeacherCourse(param)

	attendance_results := make([]map[string]interface{}, len(attendances))

	//Total student inclass
	course, _, _ := repository.QueryCourseByID(param.GetFieldValue("course_id"))
	total_query := `select count(student_id)
	from student_course_enrollments
	where course_id = ` + fmt.Sprint(course.ID)
	database.DbInstance.Raw(total_query).Scan(&total_student_inclass)

	for i, attendance := range attendances {
		student, _ := repository.QueryStudentByID(fmt.Sprint(attendance.UserID))
		user := repository.QueryUserBySID(fmt.Sprint(attendance.UserID))
		course_id := param.GetFieldValue("course_id")

		attendance_result := make(map[string]interface{})

		attendance_result["id"] = attendance.ID
		attendance_result["teacher_id"] = attendance.TeacherID
		attendance_result["student_id"] = student.StudentID
		attendance_result["student_name"] = user.LastName + " " + user.FirstName
		attendance_result["batch"] = student.Batch
		attendance_result["schedule_id"] = attendance.ScheduleID
		attendance_result["course_id"] = course_id
		attendance_result["checkin_status"] = attendance.CheckInStatus
		attendance_result["created_at"] = attendance.CreatedAt.Format("2006-01-02 15:04:05")

		attendance_result["teacher_id"] = param.GetFieldValue("teacher_id")

		switch attendance.CheckInStatus {
		case "Late":
			total_lates += 1
		case "Attend":
			total_intime += 1
		}

		attendance_results[i] = attendance_result
	}
	total_absences = len(attendances) - total_intime - total_lates

	return attendance_results, 10
}

func GetAllStudentIDs() []types.FieldOption {
	studentIds, err := repository.QueryAllStudentIDs()
	if err != nil {
		return nil
	}
	options := []types.FieldOption{}
	for _, id := range *studentIds {
		temp := types.FieldOption{
			Value: id,
			Text:  id,
		}
		options = append(options, temp)
	}
	return options
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getScheduleIds(schedules []*entity.Schedule) string {
	var data []string
	if len(schedules) > 0 {
		for _, sch := range schedules {
			f := sch.ID
			data = append(data, fmt.Sprint(f))
		}
	}
	return strings.Join(data, ",")
}
