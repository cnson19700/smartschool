package excel

import (
	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"strconv"
)

const COURSE_SHEET_NAME = "Course"

func ImportCourse(c *gin.Context) {
	w := c.Writer

	excel, err := PreprocessImport(c)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	rows := excel.GetRows(COURSE_SHEET_NAME)
	courses := make([]entity.Course, 0)
	for i := 1; i < len(rows); i += 1 {
		row := rows[i]

		var course entity.Course

		id, err := strconv.Atoi(row[0])
		if err != nil {
			w.Write([]byte("ID needs to be a number"))
			return
		}

		course.ID = uint(id)

		course.CourseID = row[1]
		course.Name = row[2]

		courses = append(courses, course)
	}

	database.DbInstance.Create(&courses)
}
