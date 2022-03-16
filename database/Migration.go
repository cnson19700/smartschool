package database

import (
	"fmt"

	"github.com/smartschool/helper"
	"github.com/smartschool/model/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func Init() {
	ConnectDatabase()
	//MigrateDatabase()
	//createDummy()
	//readDummy()
	// Close()
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
	DbInstance.AutoMigrate(&entity.Faculty{})
	DbInstance.AutoMigrate(&entity.Semester{})
	DbInstance.AutoMigrate(&entity.Role{})
	DbInstance.AutoMigrate(&entity.User{})
	DbInstance.AutoMigrate(&entity.Student{})
	DbInstance.AutoMigrate(&entity.Course{})
	DbInstance.AutoMigrate(&entity.Room{})
	DbInstance.AutoMigrate(&entity.Device{})
	DbInstance.AutoMigrate(&entity.StudentCourseEnrollment{})
	DbInstance.AutoMigrate(&entity.Schedule{})
	DbInstance.AutoMigrate(&entity.Attendance{})
	DbInstance.AutoMigrate(&entity.DeviceSignalLog{})

	errJoin := DbInstance.SetupJoinTable(&entity.Student{}, "Courses", &entity.StudentCourseEnrollment{})
	if errJoin != nil {
		panic(errJoin)
	}

	errJoin = DbInstance.SetupJoinTable(&entity.Course{}, "Students", &entity.StudentCourseEnrollment{})
	if errJoin != nil {
		panic(errJoin)
	}

	errJoin = DbInstance.SetupJoinTable(&entity.Room{}, "Courses", &entity.Schedule{})
	if errJoin != nil {
		panic(errJoin)
	}

	errJoin = DbInstance.SetupJoinTable(&entity.Course{}, "Rooms", &entity.Schedule{})
	if errJoin != nil {
		panic(errJoin)
	}

	fmt.Println("Migrate DB normal")
}

func createDummy() {

	DummyFaculties := []entity.Faculty{
		{ID: 1, Title: "Computer Science"},
		{ID: 2, Title: "Chemistry"},
		{ID: 3, Title: "Physic"},
	}

	DummyRoles := []entity.Role{
		{ID: 1, Title: "Student"},
		{ID: 2, Title: "Academic Section"},
	}

	dummyDOB := helper.StringToTimeUTC("2021-12-28T18:08:00+07:00")
	hashedPasswordByte, bcryptError := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if bcryptError != nil {
		panic(bcryptError)
	}
	DummyUsers := []entity.User{
		{ID: 1, Username: "18120001", Password: string(hashedPasswordByte), DateOfBirth: dummyDOB, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: false, PhoneNumber: "0123456789", Email: "huy@email.com", FirstName: "Huy", LastName: "Pham Quoc"},
		{ID: 2, Username: "18120002", Password: string(hashedPasswordByte), DateOfBirth: dummyDOB, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: false, PhoneNumber: "0123456789", Email: "thi@email.com", FirstName: "Thi", LastName: "Vo Thi Be"},
		{ID: 3, Username: "18120003", Password: string(hashedPasswordByte), DateOfBirth: dummyDOB, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: false, PhoneNumber: "0123456789", Email: "son@email.com", FirstName: "Son", LastName: "Cao Ngoc"},
		{ID: 4, Username: "18120004", Password: string(hashedPasswordByte), DateOfBirth: dummyDOB, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: false, PhoneNumber: "0123456789", Email: "vinh@email.com", FirstName: "Vinh", LastName: "Bui Xuan"},
		{ID: 5, Username: "18120005", Password: string(hashedPasswordByte), DateOfBirth: dummyDOB, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: false, PhoneNumber: "0123456789", Email: "nhan@email.com", FirstName: "Nhan", LastName: "Le Hoang"},
		{ID: 6, Username: "17120001", Password: string(hashedPasswordByte), DateOfBirth: dummyDOB, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: false, PhoneNumber: "0123456789", Email: "tri@email.com", FirstName: "Tri", LastName: "Ho Minh"},
	}

	DummyStudents := []entity.Student{
		{ID: 1, StudentID: "18120001", Batch: "18CTT1"},
		{ID: 2, StudentID: "18120002", Batch: "18CTT1"},
		{ID: 3, StudentID: "18120003", Batch: "18CTT1"},
		{ID: 4, StudentID: "18120004", Batch: "18CTT2"},
		{ID: 5, StudentID: "18120005", Batch: "18CTT2"},
		{ID: 6, StudentID: "17120001", Batch: "17CTT1"},
	}

	startSem, endSem := helper.StringToTimeUTC("2022-01-11T00:00:00+07:00"), helper.StringToTimeUTC("2022-04-11T00:00:00+07:00")
	DummySemester := []entity.Semester{{ID: 1, Title: "HK1", Year: "2022", FacultyID: 1, StartTime: startSem, EndTime: endSem}}

	DummyCourses := []entity.Course{
		{ID: 1, SemesterID: 1, NumberOfStudent: 40, CourseID: "CS001", Name: "Intro to Computer Science"},
		{ID: 2, SemesterID: 1, NumberOfStudent: 39, CourseID: "MTH001", Name: "Calculus I"},
		{ID: 3, SemesterID: 1, NumberOfStudent: 42, CourseID: "PH001", Name: "Physics"},
	}

	DummyStudentCourseEnrollment := []entity.StudentCourseEnrollment{
		{ID: 1, CourseID: 1, StudentID: 3},
		{ID: 2, CourseID: 1, StudentID: 4},
		{ID: 3, CourseID: 1, StudentID: 5},
		{ID: 4, CourseID: 1, StudentID: 6},
		{ID: 5, CourseID: 2, StudentID: 1},
		{ID: 6, CourseID: 2, StudentID: 2},
		{ID: 7, CourseID: 2, StudentID: 5},
		{ID: 8, CourseID: 2, StudentID: 6},
		{ID: 9, CourseID: 3, StudentID: 1},
		{ID: 10, CourseID: 3, StudentID: 2},
		{ID: 11, CourseID: 3, StudentID: 3},
		{ID: 12, CourseID: 3, StudentID: 4},
	}

	DummyRooms := []entity.Room{
		{ID: 1, RoomID: "I41", Name: "APCS lecture Room"},
		{ID: 2, RoomID: "B11", Name: "Physics lecture room"},
		{ID: 3, RoomID: "A12", Name: "Math lecture room"},
	}

	startCS, endCS := helper.StringToTimeUTC("2022-01-11T07:30:00+07:00"), helper.StringToTimeUTC("2022-01-11T09:10:00+07:00")
	startPH, endPH := helper.StringToTimeUTC("2022-01-11T09:30:00+07:00"), helper.StringToTimeUTC("2022-01-11T11:10:00+07:00")
	startCSLab, endCSLab := helper.StringToTimeUTC("2022-01-12T07:30:00+07:00"), helper.StringToTimeUTC("2022-01-12T09:30:00+07:00")
	startMTH, endMTH := helper.StringToTimeUTC("2022-01-12T09:30:00+07:00"), helper.StringToTimeUTC("2022-01-12T11:10:00+07:00")
	DummySchedule := []entity.Schedule{
		{ID: 1, RoomID: 1, CourseID: 1, StartTime: startCS, EndTime: endCS},
		{ID: 2, RoomID: 2, CourseID: 3, StartTime: startPH, EndTime: endPH},
		{ID: 3, RoomID: 3, CourseID: 2, StartTime: startMTH, EndTime: endMTH},
		{ID: 4, RoomID: 3, CourseID: 1, StartTime: startCSLab, EndTime: endCSLab},
	}

	DummyDevice := []entity.Device{
		{ID: 1, RoomID: 1, DeviceID: "D1"},
		{ID: 2, RoomID: 2, DeviceID: "D2"},
		{ID: 3, RoomID: 3, DeviceID: "D3"}}

	if DbInstance == nil {
		panic("[ERROR] Nil DB")
	}

	DbInstance.Create(&DummyFaculties)
	DbInstance.Create(&DummySemester)
	DbInstance.Create(&DummyRoles)
	DbInstance.Create(&DummyUsers)
	DbInstance.Create(&DummyStudents)
	DbInstance.Create(&DummyCourses)
	DbInstance.Create(&DummyRooms)
	DbInstance.Create(&DummyStudentCourseEnrollment)
	DbInstance.Create(&DummySchedule)
	DbInstance.Create(&DummyDevice)

	fmt.Println("Create DB dummies normal")
}

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
