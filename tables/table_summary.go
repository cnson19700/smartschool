package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/smartschool/database"
)

func GetSummary(ctx *context.Context) table.Table {
	tableSummary := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableSummary.GetInfo()

	info.AddField("ID", "id", db.Int).FieldHide()
	info.AddField("Student", "student", db.Varchar)
	info.AddField("Course", "course", db.Varchar)
	info.AddField("Absences", "absences", db.Int)

	info.SetTable("summary").SetTitle("Summary").SetDescription("Summary").
		SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
			query := `select st.name as student, c.name as course, count(*) as absences
				from attendances a, schedulers sch, students st, courses c
				where check_in_status='late' and a.scheduler_id=sch.id and a.student_id=st.id and sch.course_id=c.id
				group by st.name, c.name;`

			var summaryResults []summaryResult
			database.DbInstance.Raw(query).Scan(&summaryResults)

			tableResults := make([]map[string]interface{}, len(summaryResults))
			for i, currentResult := range summaryResults {
				tempResult := make(map[string]interface{})

				tempResult["id"] = i
				tempResult["student"] = currentResult.Student
				tempResult["course"] = currentResult.Course
				tempResult["absences"] = currentResult.Absences

				tableResults[i] = tempResult
			}

			return tableResults, 10
		})

	return tableSummary
}

type summaryResult struct {
	Student  string
	Course   string
	Absences int
}
