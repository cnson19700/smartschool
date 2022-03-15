package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/smartschool/tables"
	"net/http"
	"strconv"
	"time"
)

const SUMMARY_SHEET_NAME = "Summary"

func ExportSummary(c *gin.Context) {
	currentTime := time.Now().Unix()

	r := c.Request
	defer r.Body.Close()
	w := c.Writer

	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("Cannot parse form"))
		return
	}

	monthStr := r.Form.Get("month")
	fmt.Println(monthStr)

	excel := excelize.NewFile()

	excel.NewSheet(SUMMARY_SHEET_NAME)
	excel.DeleteSheet("Sheet1")

	excel.SetSheetRow(SUMMARY_SHEET_NAME, "A1", &[]string{"id", "student", "course", "absences"})

	data, _ := tables.GetSummaryData()
	for i, currentData := range data {
		row := make([]interface{}, 4)

		row[0] = currentData["id"]
		row[1] = currentData["student"]
		row[2] = currentData["course"]
		row[3] = currentData["absences"]

		excel.SetSheetRow(SUMMARY_SHEET_NAME, "A"+strconv.Itoa(i+2), &row)
	}

	fileName := fmt.Sprintf("Summary_%v.xlsx", currentTime)
	filePath := "public/summary_export/" + fileName

	err = excel.SaveAs(filePath)
	if err != nil {
		w.Write([]byte("Internal server error"))
		return
	}

	http.Redirect(w, r, filePath, http.StatusSeeOther)
}
