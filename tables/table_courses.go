package tables

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
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
	"gorm.io/gorm/clause"
)

func GetCourses(ctx *context.Context) table.Table {
	tableCourses := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite").SetPrimaryKey("course_id", db.Varchar))

	info := tableCourses.GetInfo()
	info.HideQueryInfo()
	info.AddCSS("span>.btn-group{display: none;}")

	info.AddField("ID", "id", db.Int).FieldSortable()
	// info.AddField("Class", "class", db.Varchar)
	info.AddField("Course Code", "course_id", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Semester", "semester_name", db.Varchar)

	info.AddColumnButtons("Attendance", types.GetColumnButton("View Attendance", icon.Info,
		action.Jump("/admin/info/attendances?course_id={%id}"),
	))

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetAllCoursesData(param)
	})
	info.SetDeleteFn(func(ids []string) error {
		for _, id := range ids {
			if len(id) != 0 {
				var dbCourse *entity.Course
				dbCourse, _, _ = repository.QueryCourseByID(id)

				if err := database.DbInstance.Delete(&dbCourse).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	info.AddButton("Delete courses", icon.Minus, action.PopUp("/delete", "Are you sure to delete courses in this semester",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
			<div>
				<form id="clean-import" method="POST" action="/delete-course-in-semester">
					<center>
						<button type="submit" class="btn btn-primary">Confirm</button>
					</center>
				</form>
			</div>`

			return true, "", data
		}))
	info.AddButton("Import Users", icon.FileExcelO, action.PopUp("/user", "Import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
				<div>
					<form id="form-import-excel" method="POST" action="/user" enctype="multipart/form-data">
						<input type="file" name="excel-file" id="file" accept=".xlsx" />
						<center>
							<input type="submit" value="????ng t???i"/>
						<center>
					</form>
				</div>`

			return true, "", data
		}))
	info.AddButton("Import courses", icon.FileExcelO, action.PopUp("/admin/info/courses", "Import",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = `
				<div>
					<form id="form-import-excel" method="POST" action="/admin/info/courses" enctype="multipart/form-data">
						<input type="file" name="excel-file" id="file" accept=".xlsx" />
						<center>
							<input type="submit" value="????ng t???i"/>
						<center>
					</form>
				</div>`

			return true, "", data
		}))
	info.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	formList := tableCourses.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("CourseID", "course_id", db.Varchar, form.Text)
	formList.AddField("Number of Students", "number_of_student", db.Int, form.Text).FieldDisplay(func(value types.FieldModel) interface{} {
		number_of_student, _ := value.Row["number_of_student"].(int)
		return number_of_student
	})
	formList.AddField("Teacher", "teacher_id", db.Varchar, form.SelectSingle).FieldOptions(GetAllTeachersData())
	formList.AddField("Teacher Role", "teacher_role", db.Varchar, form.SelectSingle).FieldOptions(types.FieldOptions{
		{
			Text:  "Professor",
			Value: "Professor",
		}, {
			Text:  "Teaching Assistant",
			Value: "Teaching Assistant",
		},
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		role, _ := value.Row["teacher_role"].(string)
		return []string{role}
	})
	formList.AddField("Semester", "semester_id", db.Int, form.SelectSingle).FieldOptions(GetAllSemestersData()).FieldDisplay(func(value types.FieldModel) interface{} {
		semester, _ := value.Row["semester_name"].(string)
		return []string{semester}
	})

	formList.AddField("Class", "class", db.Varchar, form.Text)
	// formList.EnableAjax("Success", "Fail")
	formList.SetTable("courses").SetTitle("Courses").SetDescription("Courses")

	formList.SetUpdateFn(func(values form2.Values) error {
		if values.IsEmpty("name") {
			return errors.New("Name cannot be empty")
		}
		if values.IsEmpty("course_id") {
			return errors.New("CourseID cannot be empty")
		}
		if values.IsEmpty("teacher_id") {
			return errors.New("Teacher cannot be empty")
		}
		if values.IsEmpty("teacher_role") {
			return errors.New("Teacher Role cannot be empty")
		}
		if values.IsEmpty("semester_id") {
			return errors.New("Semester cannot be empty")
		}
		if values.IsEmpty("class_id") {
			return errors.New("Class cannot be empty")
		}

		id, _ := strconv.Atoi(values.Get("id"))

		updated := database.DbInstance.Model(&entity.Course{}).Where("id = ? ", id).Clauses(clause.Returning{}).Updates(map[string]interface{}{
			"name":              values.Get("name"),
			"semester_id":       values.Get("semester_id"),
			"course_id":         values.Get("course_id"),
			"teacher_id":        values.Get("teacher_id"),
			"teacher_role":      values.Get("teacher_role"),
			"number_of_student": values.Get("number_of_student"),
			"class":             values.Get("class"),
		}).Error
		if updated != nil {
			return updated
		}
		return nil
	})

	formList.SetInsertFn(func(values form2.Values) error {
		if values.IsEmpty("name") {
			return errors.New("Name cannot be empty")
		}
		if values.IsEmpty("course_id") {
			return errors.New("CourseID cannot be empty")
		}
		if values.IsEmpty("teacher_id") {
			return errors.New("Teacher cannot be empty")
		}
		if values.IsEmpty("teacher_role") {
			return errors.New("Teacher Role cannot be empty")
		}
		if values.IsEmpty("semester_id") {
			return errors.New("Semester cannot be empty")
		}
		if values.IsEmpty("class") {
			return errors.New("Class cannot be empty")
		}
		semester_id, _ := strconv.Atoi(values.Get("semester_id"))
		num_student, _ := strconv.Atoi(values.Get("number_of_student"))
		teacher_id, _ := repository.GetIDFromTeacherID(values.Get("teacher_id"))

		if num_student < 0 {
			return errors.New("Number of student must be positive number")
		}

		course := entity.Course{
			Name:            values.Get("name"),
			SemesterID:      uint(semester_id),
			CourseID:        values.Get("course_id"),
			TeacherID:       values.Get("teacher_id"),
			TeacherRole:     values.Get("teacher_role"),
			NumberOfStudent: num_student,
			Class:           values.Get("class")}

		created := database.DbInstance.Create(&course).Error
		if created != nil {
			return created
		}
		teacher_course := entity.TeacherCourse{
			TeacherID:   teacher_id,
			TeacherRole: values.Get("teacher_role"),
			CourseID:    course.ID,
		}
		tc_created := database.DbInstance.Create(&teacher_course).Error
		if tc_created != nil {
			return tc_created
		}
		return nil
	})

	detail := tableCourses.GetDetail()
	detail.AddField("Name", "name", db.Varchar)
	detail.AddField("Semester", "semester_name", db.Varchar)
	detail.AddField("Teachers", "id", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		query := `select distinct concat(u.first_name,'', u.last_name) as teacher_name, u.id as teacher_id
		from users u
		join teacher_courses tc
		on tc.teacher_id = u.id and u.role_id = 3 and tc.course_id in (` + value.Row["id"].(string) + `)
		order by u.id`
		course_code, _ := value.Row["course_id"].(string)
		var teachers []teacherOptionResult
		database.DbInstance.Raw(query).Scan(&teachers)
		var display []interface{}
		for _, t := range teachers {
			tmp := template.Default().
				Link().
				SetURL("/admin/info/attendances?teacher_id=" + t.TeacherID + "&course_id=" + course_code).
				SetContent(template.HTML(t.TeacherName)).
				GetContent()
			display = append(display, tmp)
		}
		return display_teachers(display)

	})

	detail.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetCourseData(param.GetFieldValue(parameter.PrimaryKey))
	})

	return tableCourses
}

func GetCourseData(param string) ([]map[string]interface{}, int) {
	query := `
	select distinct u.id, u.course_id, u.name as course_name, concat(u.semester_name,' ',u.year) as semester_name, u.semester_id as semester_id
	from teacher_courses tc
	join (select c.id, c.name, c.course_id, s.title as semester_name, s.id as semester_id, s.year, c.class
		from courses c, semesters s
		where c.course_id LIKE '%` + param + `%' and c.semester_id = s.id) u
	on u.id = tc.course_id`

	var currentResult []courseResult
	database.DbInstance.Raw(query).Scan(&currentResult)
	tableResult := make([]map[string]interface{}, 1)
	tempResult := make(map[string]interface{})
	course_ids, course_names := []string{}, []string{}
	for _, c := range currentResult {
		course_ids = append(course_ids, fmt.Sprint(c.ID))
		course_names = append(course_names, c.CourseName)
	}
	tempResult["id"] = strings.Join(course_ids, ",")
	tempResult["course_id"] = currentResult[0].CourseID
	tempResult["name"] = unique_name(course_names)
	tempResult["semester_name"] = currentResult[0].SemesterName
	tempResult["semester_id"] = currentResult[0].SemesterID

	tableResult[0] = tempResult

	return tableResult, 1

}
func GetAllCoursesData(param parameter.Parameters) ([]map[string]interface{}, int) {
	query := `
	select distinct c.name as course_name, c.course_id as course_id, 
	concat(s.title,' ',s.year) as semester_name, s.id as semester_id, s.year, c.number_of_student
	from courses c, semesters s
	where c.semester_id = s.id and c.deleted_at is null order by c.course_id`

	var courseResults []courseResult
	database.DbInstance.Raw(query).Scan(&courseResults)
	tableResults := make([]map[string]interface{}, len(courseResults))
	for i, currentResult := range courseResults {
		tempResult := make(map[string]interface{})

		tempResult["id"] = i + 1
		// tempResult["teacher_name"] = currentResult.TeacherName
		// tempResult["teacher_id"] = currentResult.TeacherID
		tempResult["course_id"] = currentResult.CourseID
		tempResult["name"] = currentResult.CourseName
		tempResult["semester_name"] = currentResult.SemesterName
		// tempResult["number_of_student"] = currentResult.NumberOfStudent
		tempResult["semester_id"] = currentResult.SemesterID
		// tempResult["teacher_role"] = currentResult.TeacherRole
		// tempResult["class"] = currentResult.Class

		tableResults[i] = tempResult
	}
	return tableResults, 10
}

func GetAllTeachersData() []types.FieldOption {
	teacher_options := []types.FieldOption{}

	teacher_results := []teacherOptionResult{}

	query := `select t.teacher_id, concat(u.first_name, ' ', u.last_name) as teacher_name
				from users u, teachers t
				where u.id = t.id	
		`

	database.DbInstance.Raw(query).Scan(&teacher_results)
	for _, teacher := range teacher_results {
		tmp := types.FieldOption{
			Text:  teacher.TeacherName,
			Value: teacher.TeacherID,
		}
		teacher_options = append(teacher_options, tmp)
	}
	return teacher_options
}

func GetAllSemestersData() []types.FieldOption {
	semester_options := []types.FieldOption{}

	semester_results := []semesterOptionResult{}

	query := `select s.id, s.title
				from semesters s`

	database.DbInstance.Raw(query).Scan(&semester_results)
	for _, semester := range semester_results {
		tmp := types.FieldOption{
			Text:  semester.Title,
			Value: fmt.Sprint(semester.ID),
		}
		semester_options = append(semester_options, tmp)
	}
	return semester_options
}

func display_teachers(og []interface{}) (display string) {
	data := make([]string, len(og))
	for i, v := range og {
		data[i] = fmt.Sprint(v)
	}
	return strings.Join(data, "<br>")
}
func unique_name(s []string) string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return strings.Join(result, "/")
}

type courseResult struct {
	ID              int
	TeacherName     string
	TeacherID       string
	NumberOfStudent int
	CourseName      string
	CourseID        string
	SemesterName    string
	SemesterID      int
	TeacherRole     string
	Class           string
}

type teacherOptionResult struct {
	TeacherID   string
	TeacherName string
}

type semesterOptionResult struct {
	ID    int
	Title string
}
