package database

import (
	"fmt"
	"time"

	"github.com/smartschool/helper"
	"github.com/smartschool/model/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func Init() {
	ConnectDatabase()
	// DropAllTables()
	// MigrateDatabase()
	// AttendanceData()
	// createDummy2()
	//readDummy()
	// Close()
}

func Close() {
	safe, _ := DbInstance.DB()

	defer safe.Close()
}

func ConnectDatabase() {
	dbURI := "host=13.228.244.196 port=5432 user=busmapdb dbname=phenikaamaas_attendancedb sslmode=disable password=frjsdfhaflpzlcdzgnfvuxkdwiiiiklpojzowxajmendeeoqtbzyrgi"

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
	DbInstance.AutoMigrate(&entity.Teacher{})
	DbInstance.AutoMigrate(&entity.Room{})
	DbInstance.AutoMigrate(&entity.Device{})
	DbInstance.AutoMigrate(&entity.StudentCourseEnrollment{})
	DbInstance.AutoMigrate(&entity.Schedule{})
	DbInstance.AutoMigrate(&entity.TeacherCourse{})
	DbInstance.AutoMigrate(&entity.Attendance{})
	DbInstance.AutoMigrate(&entity.DeviceSignalLog{})
	DbInstance.AutoMigrate(&entity.UserNotificationToken{})

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

	//errJoin = DbInstance.SetupJoinTable(&entity.Teacher{}, "Courses", &entity.TeacherCourse{})
	//if errJoin != nil {
	//	panic(errJoin)
	//}

	fmt.Println("Migrate DB normal")
}

func addScheduleDummy() {
	startSem, _ := helper.StringToTimeUTC("2022-04-11T00:00:00+07:00")
	endSem, _ := helper.StringToTimeUTC("2022-09-11T00:00:00+07:00")
	DummySemester := []entity.Semester{{ID: 2, Title: "HK2", Year: "2022", FacultyID: 1, StartTime: startSem, EndTime: endSem}}

	DummyCourses := []entity.Course{
		{ID: 4, SemesterID: 2, NumberOfStudent: 45, CourseID: "CS002", Name: "Intro to Computer Science II"},
		{ID: 5, SemesterID: 2, NumberOfStudent: 44, CourseID: "MTH002", Name: "Calculus II"},
		{ID: 6, SemesterID: 2, NumberOfStudent: 45, CourseID: "PH002", Name: "Physics II"},
	}

	startCS, _ := helper.StringToTimeUTC("2022-04-12T07:30:00+07:00")
	endCS, _ := helper.StringToTimeUTC("2022-04-12T09:10:00+07:00")
	startPH, _ := helper.StringToTimeUTC("2022-04-12T09:30:00+07:00")
	endPH, _ := helper.StringToTimeUTC("2022-04-12T11:10:00+07:00")
	startCSLab, _ := helper.StringToTimeUTC("2022-04-12T13:30:00+07:00")
	endCSLab, _ := helper.StringToTimeUTC("2022-04-12T15:30:00+07:00")
	startMTH, _ := helper.StringToTimeUTC("2022-04-12T15:30:00+07:00")
	endMTH, _ := helper.StringToTimeUTC("2022-04-12T17:10:00+07:00")
	DummySchedule := []entity.Schedule{
		{ID: 5, RoomID: 1, CourseID: 4, StartTime: startCS, EndTime: endCS},
		{ID: 6, RoomID: 2, CourseID: 5, StartTime: startPH, EndTime: endPH},
		{ID: 7, RoomID: 3, CourseID: 6, StartTime: startMTH, EndTime: endMTH},
		{ID: 8, RoomID: 4, CourseID: 5, StartTime: startCSLab, EndTime: endCSLab},
	}

	DummyStudentCourseEnrollment := []entity.StudentCourseEnrollment{
		{ID: 13, CourseID: 4, StudentID: 3},
		{ID: 14, CourseID: 4, StudentID: 4},
		{ID: 15, CourseID: 4, StudentID: 5},
		{ID: 16, CourseID: 4, StudentID: 6},
		{ID: 17, CourseID: 5, StudentID: 1},
		{ID: 18, CourseID: 5, StudentID: 2},
		{ID: 19, CourseID: 5, StudentID: 5},
		{ID: 20, CourseID: 5, StudentID: 6},
		{ID: 21, CourseID: 6, StudentID: 1},
		{ID: 22, CourseID: 6, StudentID: 2},
		{ID: 23, CourseID: 6, StudentID: 3},
		{ID: 24, CourseID: 6, StudentID: 4},
	}

	DbInstance.Create(&DummySemester)
	DbInstance.Create(&DummyCourses)
	DbInstance.Create(&DummyStudentCourseEnrollment)
	DbInstance.Create(&DummySchedule)

}

func createDummy2() {

	DummyFaculties := []entity.Faculty{
		{ID: 1, Title: "Computer Science"},
		{ID: 2, Title: "Chemistry"},
		{ID: 3, Title: "Physic"},
	}

	DummyRoles := []entity.Role{
		{ID: 1, Title: "Student"},
		{ID: 2, Title: "Academic Section"},
		{ID: 3, Title: "Professor"},
	}

	dummyDOB, _ := helper.StringToTimeUTC("2021-12-28T18:08:00+07:00")
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
		{ID: 7, Username: "Dinh Ba Tien", Password: string(hashedPasswordByte), Email: "dbtien@capstone.local", PhoneNumber: "0123456789", FirstName: "Dinh", LastName: "Tien", DateOfBirth: dummyDOB, RoleID: 3, Gender: 0, FacultyID: 1, IsActivate: true, Teacher: &entity.Teacher{TeacherID: "10001"}},
		{ID: 8, Username: "Nhat Hoa", Password: string(hashedPasswordByte), Email: "nhoa@capstone.local", PhoneNumber: "0123456789", FirstName: "Nhat", LastName: "Hoa", DateOfBirth: dummyDOB, RoleID: 3, Gender: 1, FacultyID: 3, IsActivate: true, Teacher: &entity.Teacher{TeacherID: "20010"}},
		{ID: 9, Username: "Ngoc Hue", Password: string(hashedPasswordByte), Email: "nghue@capstone.local", PhoneNumber: "0123456789", FirstName: "Ngoc", LastName: "Hue", DateOfBirth: dummyDOB, RoleID: 3, Gender: 1, FacultyID: 2, IsActivate: true, Teacher: &entity.Teacher{TeacherID: "30011"}},
	}

	DummyStudents := []entity.Student{
		{ID: 1, StudentID: "18120001", Batch: "18CTT1"},
		{ID: 2, StudentID: "18120002", Batch: "18CTT1"},
		{ID: 3, StudentID: "18120003", Batch: "18CTT1"},
		{ID: 4, StudentID: "18120004", Batch: "18CTT2"},
		{ID: 5, StudentID: "18120005", Batch: "18CTT2"},
		{ID: 6, StudentID: "17120001", Batch: "17CTT1"},
	}

	startSem, _ := helper.StringToTimeUTC("2022-01-11T00:00:00+07:00")
	endSem, _ := helper.StringToTimeUTC("2022-04-11T00:00:00+07:00")
	DummySemester := []entity.Semester{{ID: 1, Title: "HK1", Year: "2022", FacultyID: 1, StartTime: startSem, EndTime: endSem}}

	DummyCourses := []entity.Course{
		{ID: 1, SemesterID: 1, NumberOfStudent: 40, CourseID: "CS001", Name: "Intro to Computer Science", TeacherID: "10001", TeacherRole: "Professor"},
		{ID: 2, SemesterID: 1, NumberOfStudent: 39, CourseID: "MTH001", Name: "Calculus I", TeacherID: "20010", TeacherRole: "Professor"},
		{ID: 3, SemesterID: 1, NumberOfStudent: 42, CourseID: "PH001", Name: "Physics", TeacherID: "30011", TeacherRole: "Professor"},
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

	startCS, _ := helper.StringToTimeUTC("2022-01-11T07:30:00+07:00")
	endCS, _ := helper.StringToTimeUTC("2022-01-11T09:10:00+07:00")
	startPH, _ := helper.StringToTimeUTC("2022-01-11T09:30:00+07:00")
	endPH, _ := helper.StringToTimeUTC("2022-01-11T11:10:00+07:00")
	startCSLab, _ := helper.StringToTimeUTC("2022-01-12T07:30:00+07:00")
	endCSLab, _ := helper.StringToTimeUTC("2022-01-12T09:30:00+07:00")
	startMTH, _ := helper.StringToTimeUTC("2022-01-12T09:30:00+07:00")
	endMTH, _ := helper.StringToTimeUTC("2022-01-12T11:10:00+07:00")
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

	DummyTeacherCourses := []entity.TeacherCourse{{TeacherID: 7, CourseID: 1}, {TeacherID: 8, CourseID: 3}}

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
	DbInstance.Create(&DummyTeacherCourses)

	fmt.Println("Create DB dummies normal")
}

func createDummy() {
	t := time.Now()
	DummyFaculties := []entity.Faculty{{Title: "Computer Science"}, {Title: "Chemistry"}}
	DummySemesters := []entity.Semester{{Title: "HK1-2018", Year: "2018", StartTime: time.Now(), EndTime: time.Now().Add(20), FacultyID: 1}}
	DummyRoles := []entity.Role{{Title: "Student"}, {Title: "Academic Section"}, {Title: "Teacher"}}
	DummyUsers := []entity.User{
		{Username: "Bui Xuan Vinh", Password: "12345", Email: "vinh@capstone.local", PhoneNumber: "0123456789", FirstName: "Bui", LastName: "Vinh", DateOfBirth: t, RoleID: 1, Gender: 0, FacultyID: 1, IsActivate: true, Student: &entity.Student{StudentID: "100", Batch: "18CTT2"}},
		{Username: "Dinh Ba Tien", Password: "12345", Email: "dbtien@capstone.local", PhoneNumber: "0123456789", FirstName: "Dinh", LastName: "Ba", DateOfBirth: t, RoleID: 3, Gender: 0, FacultyID: 1, IsActivate: true, Teacher: &entity.Teacher{TeacherID: "10001"}},
	}
	DummyCourses := []entity.Course{{CourseID: "CS001", Name: "Intro to Internet", SemesterID: 2}, {CourseID: "MTH001", Name: "Intro to Math", SemesterID: 2}}
	DummyRooms := []entity.Room{{RoomID: "I41", Name: "APCS Room"}, {RoomID: "B52", Name: "CLC lab"}, {RoomID: "E15", Name: "VP Stone room"}}
	DummyDevice := []entity.Device{{RoomID: 1, DeviceID: "D1"}, {RoomID: 2, DeviceID: "D2"}, {RoomID: 3, DeviceID: "D3"}}
	DummyTeacherCourses := []entity.TeacherCourse{{TeacherID: 2, CourseID: 9}}
	if DbInstance == nil {
		panic("[ERROR] Nil DB")
	}

	DbInstance.Create(&DummyFaculties)
	DbInstance.Create(&DummySemesters)
	DbInstance.Create(&DummyRoles)
	DbInstance.Create(&DummyUsers)
	DbInstance.Create(&DummyCourses)
	DbInstance.Create(&DummyRooms)
	DbInstance.Create(&DummyDevice)
	DbInstance.Create(&DummyTeacherCourses)

}

func DropAllTables() {
	DbInstance.Migrator().DropTable(&entity.Faculty{})
	DbInstance.Migrator().DropTable(&entity.Semester{})
	DbInstance.Migrator().DropTable(&entity.Role{})
	DbInstance.Migrator().DropTable(&entity.User{})
	DbInstance.Migrator().DropTable(&entity.Student{})
	DbInstance.Migrator().DropTable(&entity.Course{})
	DbInstance.Migrator().DropTable(&entity.Teacher{})
	DbInstance.Migrator().DropTable(&entity.Room{})
	DbInstance.Migrator().DropTable(&entity.Device{})
	DbInstance.Migrator().DropTable(&entity.StudentCourseEnrollment{})
	DbInstance.Migrator().DropTable(&entity.Schedule{})
	DbInstance.Migrator().DropTable(&entity.Attendance{})
	DbInstance.Migrator().DropTable(&entity.DeviceSignalLog{})
	DbInstance.Migrator().DropTable(&entity.TeacherCourse{})
}

func AttendanceData() {
	DummyAttendances := []entity.Attendance{
		{ID: 1, UserID: 1, TeacherID: 10001, ScheduleID: 1, CheckInTime: time.Now(), CheckInStatus: "Attend"},
		{ID: 2, UserID: 2, TeacherID: 10001, ScheduleID: 1, CheckInTime: time.Now().Add(1000), CheckInStatus: "Late"},
		{ID: 3, UserID: 5, TeacherID: 10001, ScheduleID: 1, CheckInTime: time.Now().Add(3000), CheckInStatus: "Late"},
	}

	DbInstance.Create(&DummyAttendances)
}
