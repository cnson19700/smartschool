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
	"github.com/smartschool/helper"
)

func GetSummary(ctx *context.Context) table.Table {
	tableSummary := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := tableSummary.GetInfo()

	info.AddField("ID", "id", db.Int).FieldHide()
	info.AddField("Student", "student", db.Varchar)
	info.AddField("Course", "course", db.Varchar)
	info.AddField("Absences", "absences", db.Int)

	info.SetGetDataFn(func(param parameter.Parameters) ([]map[string]interface{}, int) {
		return GetSummaryData()
	})

	info.AddButton("Export summary", icon.FileExcelO, action.PopUp("/summary", "Export",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			data = fmt.Sprintf(`
				<form id="form-export-summary" method="GET" action="/summary">
	        		<div class="form-group" style="display: flex; align-items: center">
                		<label for="month" class="control-label" style="width:150px">Chọn tháng</label>
						%v
					</div>
            		<div class="form-group" style="display:flex; justify-content:center">
                		<input class="btn btn-sm btn-primary submit" style="width:100px; height: 35px; font-size: 16px" type="submit" value="Xuất"/>
            		</div>
        		</form>`, helper.GetMonthSelectBox())

			return true, "", data
		}))

	info.SetTable("summary").SetTitle("Summary").SetDescription("Summary")

	return tableSummary
}

func GetSummaryData() ([]map[string]interface{}, int) {
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
}

type summaryResult struct {
	Student  string
	Course   string
	Absences int
}
