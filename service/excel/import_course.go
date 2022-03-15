package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const COURSE_SHEET_NAME = "Course"

func ImportCourse(c *gin.Context) {
	currentTime := time.Now().Unix()

	r := c.Request
	defer r.Body.Close()
	w := c.Writer

	file, fileHeader, err := r.FormFile("excel-file")
	if err != nil {
		w.Write([]byte("Invalid uploaded file"))
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte("Invalid uploaded file"))
		return
	}

	fileNameWithExt := fileHeader.Filename
	fileExt := filepath.Ext(fileNameWithExt)
	fileNameOnly := strings.TrimSuffix(fileNameWithExt, fileExt)
	fileNameSaved := fmt.Sprintf("%s_%d%s", fileNameOnly, currentTime, fileExt)
	filePath := "public/course_import/" + fileNameSaved

	err = ioutil.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		w.Write([]byte("Internal server error"))
		return
	}

	excel, err := excelize.OpenFile(filePath)
	if err != nil {
		w.Write([]byte("Internal server error"))
		return
	}

	rows := excel.GetRows(COURSE_SHEET_NAME)
	courses := make([]entity.Course, 0)
	for i := 1; i < len(rows); i += 1 {
		row := rows[i]

		var course entity.Course

		id, err := strconv.Atoi(row[0])
		if err != nil {
			w.Write([]byte("Cannot parse file"))
			return
		}

		course.ID = uint(id)

		course.CourseID = row[1]
		course.Name = row[2]

		courses = append(courses, course)
	}

	database.DbInstance.Create(&courses)
}
