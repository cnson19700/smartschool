package excel

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/helper"
	"github.com/smartschool/model/entity"
	"github.com/smartschool/repository"
)

type course_row struct {
	course_id  int
	row        int
	start_time time.Time
	end_time   time.Time
}

func ImportSchedule(c *gin.Context) {
	w := c.Writer
	excel, err := PreprocessImport(c, "public/schedule_import/")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	const QTTB_SHEET_NAME = "tuần 1"
	// splitRegExp := regexp.MustCompile(`\n|,|-`)
	rooms := make([]entity.Room, 0)
	courses := []course_row{}
	times := []string{}

	var tm int
	database.DbInstance.Table("rooms").Select("setval('rooms_id_seq', (SELECT MAX(id) FROM rooms))").Scan(tm)
	// Get next id for room
	f_row := 0
	for idr, rows := range excel.GetRows(QTTB_SHEET_NAME) {

		if idr == 0 {
			continue
		}
		//Get times
		if idr == 1 {
			for _, c := range rows {
				if c == "" {
					continue
				}
				t := strings.ReplaceAll(c, "h", ":")

				tmp := strings.Split(t, " - ")
				if len(tmp) > 1 {
					for _, v := range tmp {
						times = append(times, "2022-04-12T"+v+":00+07:00")
					}
				} else {
					times = append(times, "2022-04-12T"+t+":00+07:00")
				}
			}
			continue
		}

		// Get rooms
		r := rows[7]
		var room entity.Room

		room.RoomID = r
		room.Name = r

		rooms = append(rooms, room)
		// Get Course
		for j := range rows {
			if rows[j] != "" && j != 7 {
				course_code, _ := GetCourse(rows[j])
				course_id := repository.QueryCourseIndexByCode(course_code)
				courses = append(courses,
					course_row{
						course_id:  course_id,
						row:        f_row,
						start_time: parseTime(times[j]),
						end_time:   parseTime(times[j+1])})
			}
		}
		f_row += 1
	}
	database.DbInstance.Create(&rooms)

	AddSchedule(courses, rooms)
}

func GetCourse(init string) (string, string) {
	init = strings.ReplaceAll(init, "/", "(")
	course := strings.Split(init, "(")
	return strings.TrimSpace(course[0]), course[1]
}

func AddSchedule(cr []course_row, r []entity.Room) {
	var tm int
	database.DbInstance.Table("schedules").Select("setval('schedules_id_seq', (SELECT MAX(id) FROM schedules))").Scan(tm)
	schedules := make([]entity.Schedule, 0)
	for _, c := range cr {
		schedule := entity.Schedule{}

		schedule.CourseID = uint(c.course_id)
		schedule.RoomID = r[c.row].ID
		schedule.StartTime = c.start_time
		schedule.EndTime = c.end_time

		schedules = append(schedules, schedule)
	}
	database.DbInstance.Create(&schedules)
}

func parseTime(str string) time.Time {
	time, _ := helper.StringToTimeUTC(str)
	return time
}
