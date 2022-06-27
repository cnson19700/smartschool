package excel

import (
	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var COURSE_SHEET_NAMES = []string{"CLC", "CTTT", "VP", "Tổng hợp"}

func ImportCourse(c *gin.Context) {
	ImportTeacherFromCourseFile(c)

	w := c.Writer

	excel, err := PreprocessImport(c)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	splitRegExp := regexp.MustCompile(`\n|,|-`)

	for _, sheetName := range COURSE_SHEET_NAMES {
		rows := excel.GetRows(sheetName)

		for i := 5; i < len(rows); i += 1 {
			row := rows[i]

			_, notFound, err := repository.QueryCourseByCourseIdAndClass(row[1], row[2])
			if err != nil {
				w.Write([]byte("Error calling database"))
				return
			}

			if !notFound {
				continue
			}

			var course entity.Course

			course.CourseID = row[1]
			course.Class = row[2]
			course.Name = row[3]
			course.SemesterID = 1

			database.DbInstance.Create(&course)

			for j := 4; j < 7; j += 1 {
				cell := row[j]
				names := splitRegExp.Split(cell, -1)

				for _, name := range names {
					name := strings.Trim(name, " \n")

					if name == "" {
						continue
					}

					teacher, notFound, err := repository.QueryTeacherByName(name)
					if err != nil {
						w.Write([]byte("Error calling database"))
						return
					}
					if notFound {
						w.Write([]byte("Teacher not found: " + name))
						return
					}

					var teacherCourse entity.TeacherCourse

					teacherCourse.CourseID = course.ID
					teacherCourse.TeacherID = teacher.ID
					if j == 4 {
						teacherCourse.TeacherRole = "GVLT"
					} else if j == 5 {
						teacherCourse.TeacherRole = "HDTH"
					} else {
						teacherCourse.TeacherRole = "TG"
					}

					database.DbInstance.Create(&teacherCourse)
				}
			}
		}
	}

	c.Redirect(301, "/admin/info/courses")
}

func ImportTeacherFromCourseFile(c *gin.Context) {
	w := c.Writer

	excel, err := PreprocessImport(c)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	users := make([]entity.User, 0)
	splitRegExp := regexp.MustCompile(`\n|,|-`)
	nameSet := make(map[string]struct{})
	rand.Seed(time.Now().UnixNano())

	for _, sheetName := range COURSE_SHEET_NAMES {
		rows := excel.GetRows(sheetName)

		for i := 5; i < len(rows); i += 1 {
			row := rows[i]

			for j := 4; j < 7; j += 1 {
				cell := row[j]
				names := splitRegExp.Split(cell, -1)

				for _, name := range names {
					name := strings.Trim(name, " \n")

					if name == "" {
						continue
					}

					_, found := nameSet[name]
					if found {
						continue
					}
					nameSet[name] = struct{}{}

					_, notFound, err := repository.QueryTeacherByName(name)
					if err != nil {
						w.Write([]byte("Error calling database"))
						return
					}

					if !notFound {
						continue
					}

					var user entity.User

					user.Username = name
					user.FirstName = name
					user.DateOfBirth = time.Now()
					user.RoleID = 3
					user.FacultyID = 1
					user.Teacher = &entity.Teacher{
						TeacherID: strconv.Itoa(rand.Intn(100000)),
					}

					users = append(users, user)
				}
			}
		}
	}

	database.DbInstance.Create(&users)
}
