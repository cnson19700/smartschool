package database

import (
	"fmt"

	"github.com/smartschool/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func Init() {
	ConnectDatabase()
	MigrateDatabase()
	//createDummy()
	//readDummy()
}

func Close() {
	safe, _ := DbInstance.DB()

	defer safe.Close()
}

func ConnectDatabase() {
	//dbURI := "host=13.228.244.196 port=5432 user=busmapdb dbname=phenikaamaas_attendancedb sslmode=disable password=frjsdfhaflpzlcdzgnfvuxkdwiiiiklpojzowxajmendeeoqtbzyrgi"
	dbURI := "host=localhost user=postgres dbname=nhan_local_database sslmode=disable password=Postgres port=5432"

	var err error
	DbInstance, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect DB normal")
	}
}

func MigrateDatabase() {
	DbInstance.AutoMigrate(&entity.Student{})
	DbInstance.AutoMigrate(&entity.Course{})
	DbInstance.AutoMigrate(&entity.Room{})
	DbInstance.AutoMigrate(&entity.StudentCourse{})
	DbInstance.AutoMigrate(&entity.Scheduler{})
	DbInstance.AutoMigrate(&entity.Device{})
	DbInstance.AutoMigrate(&entity.Attendance{})
	DbInstance.AutoMigrate(&entity.User{})

	errJoin := DbInstance.SetupJoinTable(&entity.Student{}, "Courses", &entity.StudentCourse{})

	if errJoin != nil {
		panic(errJoin)
	}

	errJoin = DbInstance.SetupJoinTable(&entity.Room{}, "Courses", &entity.Scheduler{})

	if errJoin != nil {
		panic(errJoin)
	}

	fmt.Println("Migrate DB normal")
}

// func createDummy() {
// 	t := time.Now()
// 	DummyStudents := []entity.Student{{StudentID: "100", Name: "Alice", Email: "abc@mail.com", PhoneNumber: "123456789"}, {StudentID: "101", Name: "Bob", Email: "abc@mail.com", PhoneNumber: "123456789"}, {StudentID: "102", Name: "Carly", Email: "abc@mail.com", PhoneNumber: "123456789"}}
// 	DummyCourses := []entity.Course{{CourseID: "CS001", Name: "Intro to Internet"}, {CourseID: "MTH001", Name: "Intro to Math"}}
// 	DummyStudentCourse := []entity.StudentCourse{{StudentID: 1, CourseID: 1}, {StudentID: 1, CourseID: 2}, {StudentID: 2, CourseID: 1}, {StudentID: 3, CourseID: 2}}
// 	DummyRooms := []entity.Room{{RoomID: "I41", Name: "APCS Room"}, {RoomID: "B52", Name: "CLC lab"}, {RoomID: "E15", Name: "VP Stone room"}}
// 	DummyScheduler := []entity.Scheduler{{RoomID: 1, CourseID: 1, StartTime: t, EndTime: t.Add(time.Hour * 2)}, {RoomID: 2, CourseID: 2, StartTime: t, EndTime: t.Add(time.Hour * 2)}, {RoomID: 3, CourseID: 1, StartTime: t.Add(time.Hour * 4), EndTime: t.Add(time.Hour * 6)}, {RoomID: 1, CourseID: 2, StartTime: t.Add(time.Hour * 2), EndTime: t.Add(time.Hour * 4)}}
// 	DummyDevice := []entity.Device{{RoomID: 1, DeviceID: "D1"}, {RoomID: 2, DeviceID: "D2"}, {RoomID: 3, DeviceID: "D3"}}

// 	if DbInstance == nil {
// 		panic("[ERROR] Nil DB")
// 	}

// 	DbInstance.Create(&DummyStudents)
// 	DbInstance.Create(&DummyCourses)
// 	DbInstance.Create(&DummyStudentCourse)
// 	DbInstance.Create(&DummyRooms)
// 	DbInstance.Create(&DummyScheduler)
// 	DbInstance.Create(&DummyDevice)
// }

// func readDummy() {
// 	studentID := "100"
// 	deviceID := "D1"
// 	t0 := "2022-02-16T9:59:00Z"

// 	t, err := time.Parse(time.RFC3339, t0)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

// 	fmt.Println(t.In(loc))
// 	x := t.In(loc).UTC()
// 	fmt.Print(x)

// 	var device entity.Device
// 	DbInstance.Select("room_id").Where("device_id = ?", deviceID).Find(&device)

// 	var result entity.Scheduler
// 	DbInstance.Select("course_id", "end_time").Where("room_id = ? AND start_time <= ? AND end_time > ?", device.RoomID, t, t).Preload("Course").Find(&result)

// 	fmt.Print(result.Course.CourseID)

// 	var student entity.Student
// 	DbInstance.Select("id").Where("student_id = ?", studentID).First(&student)
// 	fmt.Println(student.ID)

// 	var verify entity.StudentCourse
// 	DbInstance.Where("student_id = ? AND course_id = ?", student.ID, result.CourseID).Find(&verify)

// 	if verify.ID != 0 {
// 		var checkAttend entity.Attendance
// 		DbInstance.Select("id").Where("student_id = ? AND course_id = ? AND room_id = ? AND end_time > ?", verify.StudentID, verify.CourseID, device.RoomID, t.In(loc)).Find(&checkAttend)

// 		if checkAttend.ID == 0 {
// 			DbInstance.Create(&entity.Attendance{StudentID: verify.StudentID, CourseID: verify.CourseID, RoomID: device.RoomID, CheckInTime: t.In(loc), EndTime: result.EndTime, CheckInStatus: "Late"})
// 		} else {
// 			fmt.Println("Checkin exist!!!")
// 		}
// 	} else {
// 		fmt.Println("Record not found")
// 	}
// }
