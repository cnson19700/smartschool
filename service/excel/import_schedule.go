package excel

import (
	"reflect"
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
	row        string
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

	var QTTB_SHEET_NAME = "Sheet1"
	// splitRegExp := regexp.MustCompile(`\n|,|-`)
	rooms := make([]string, 0)
	courses := []course_row{}
	times := []string{}

	var tm int
	// var cols int
	database.DbInstance.Table("rooms").Select("setval('rooms_id_seq', (SELECT MAX(id) FROM rooms))").Scan(tm)
	// Get next id for room
	f_row := 0

	var date string
	var num_col int
	for idr, rows := range excel.GetRows(QTTB_SHEET_NAME) {
		if idr == 0 {
			num_col = countCol(rows)
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
						times = append(times, "T"+v+":00+07:00")
					}
				} else {
					times = append(times, "T"+t+":00+07:00")
				}
			}
			continue
		}

		// Get Date
		if contains(rows, "Phòng") {
			tmp := strings.Split(rows[0], " ")
			date = strings.Join(tmp[len(tmp)-2:], "")
			tmp = strings.Split(date, "/")
			swapF := reflect.Swapper(tmp)

			for i := 0; i < len(tmp)/2; i++ {
				swapF(i, len(tmp)-1-i)
			}
			date = strings.Join(tmp, "-")
			continue
		}

		if contains(rows, "6h40") {
			continue
		}

		// Get rooms name
		r := rows[7]
		if r != "" {
			rooms = append(rooms, r)
		}

		// Get Course
		for j := range rows {
			if j >= num_col {
				continue
			}
			if rows[j] != "" && j != 7 {
				course_code, _ := GetCourse(rows[j])
				course_id := repository.QueryCourseIndexByCode(course_code)
				if course_id == 0 {
					continue
				}
				courses = append(courses,
					course_row{
						course_id:  course_id,
						row:        rows[7],
						start_time: parseTime(strings.ReplaceAll(times[j], "T", date+"T")),
						end_time:   parseTime(strings.ReplaceAll(times[j+1], "T", date+"T"))})
			}
		}
		f_row += 1
	}
	rooms = removeDuplicateStr(rooms)
	rooms_en := buildRooms(rooms)
	database.DbInstance.Create(&rooms_en)

	AddSchedule(courses, rooms_en)
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
		schedule.RoomID = repository.QueryRoomByName(c.row)
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

func countCol(rows []string) int {
	var r []string
	for _, str := range rows {
		if str != "" {
			r = append(r, str)
		}
	}
	return len(r)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func buildRooms(str []string) (res []entity.Room) {
	for _, room := range str {
		var r entity.Room
		r.RoomID = room
		r.Name = room
		res = append(res, r)
	}
	return res
}
