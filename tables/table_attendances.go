package tables

import (
	"fmt"

	template2 "html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

var total_lates, total_intime, total_absences, total_student_inclass = 0, 0, 0, 0
var batch string

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
	}).HideFilterButton()
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
		switch c {
		case "Late":
			c = "<span id='late_stt'>" + c + "</span>"
		case "Attend":
			c = "<span id='attend-stt'>" + c + "</span>"
		}
		return c
	})
	info.AddField("Created At", "created_at", db.Timestamp).FieldFilterable(types.FilterType{FormType: form.DateRange, Placeholder: " ... "}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			c, _ := value.Row["created_at"].(string)
			return "<span id='att-created-at'>" + c + "</span>"
		})
	info.HideNewButton()
	info.HideDetailButton()
	info.HideDeleteButton()
	info.HideQueryInfo()
	info.AddCSS(".late-stt{color: red;} .attend-stt{color: rgb(23, 246, 67);}")

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
<<<<<<< HEAD
				{"Title": types.InfoItem{Content: "Lates"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_lates))}},
				{"Title": types.InfoItem{Content: "Attend"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_intime))}},
				{"Title": types.InfoItem{Content: "Absences"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_absences))}},
=======
				{"Status": types.InfoItem{Content: "Lates"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_lates) + "/" + fmt.Sprint(total_student_inclass))}},
				{"Status": types.InfoItem{Content: "In times"},
					"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_intime) + "/" + fmt.Sprint(total_student_inclass))}},
				// {"Title": types.InfoItem{Content: "Absences"},
				// 	"Times": types.InfoItem{Content: template2.HTML(fmt.Sprint(total_absences))}},
>>>>>>> 9c0a08f (update v1)
			},
			).SetMinWidth("100px").GetContent()
			col2 := `<div style="position: absolute;width:230px;">` + template.Default().Box().SetHeader("Overview").
				WithHeadBorder().SetBody(table_sum).GetContent() + `</div>`
			return `<div style="width:100%;">` + col2 + col1 + `</div>`
		})
	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetAllAttendancesData(param) //base on teacher_course
	})
	info.SetTable("attendances").SetTitle("Attendances " + ctx.FormValue("course_id") + " - " + batch).SetDescription("Attendances")
	return tableAttendaces
}

func GetAllAttendancesData(param parameter.Parameters) ([]map[string]interface{}, int) {
	total_absences, total_intime, total_lates = 0, 0, 0
	attendances := []*entity.Attendance{}
	attendances, _ = repository.AttendanceByTeacherCourse(param)

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
		attendance_result["student_name"] = user.FirstName + " " + user.LastName
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
