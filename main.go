package main

import (
	"fmt"

	"entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbEntity *gorm.DB
var err error

func createDummy() {
	DummyStudents := []entity.Student{{StudentID: "100", Name: "Alice", Email: "abc@mail.com", PhoneNumber: "123456789"}, {StudentID: "101", Name: "Bob", Email: "abc@mail.com", PhoneNumber: "123456789"}, {StudentID: "102", Name: "Carly", Email: "abc@mail.com", PhoneNumber: "123456789"}}
	DummyCourses := []entity.Course{{CourseID: "CS001", Name: "Intro to Internet"}, {CourseID: "MTH001", Name: "Intro to Math"}}
	DummyStudentCourse := []entity.StudentCourse{{StudentID: 1, CourseID: 1}, {StudentID: 1, CourseID: 2}, {StudentID: 2, CourseID: 1}, {StudentID: 3, CourseID: 2}}

	if DbEntity == nil {
		panic("[ERROR] Nil DB")
	}

	DbEntity.Create(&DummyStudents)
	DbEntity.Create(&DummyCourses)
	DbEntity.Create(&DummyStudentCourse)
}

func readDummy() {
	// result := map[string]interface{}{}
	// db.Table("students").Take(&result)
	fmt.Println("=========Student=========")

	var stus []entity.Student
	DbEntity.Preload("Courses").Find(&stus)
	//teas := db.Preload("Students").Limit(2).Find(&Teacher{})

	for i := 0; i < len(stus); i++ {
		fmt.Println(stus[i])
	}
	fmt.Println("=========Courses=========")
	var cous []entity.Course
	DbEntity.Preload("Students").Find(&cous)
	//teas := db.Preload("Students").Limit(2).Find(&Teacher{})

	for i := 0; i < len(cous); i++ {
		fmt.Println(cous[i])

		for j := 0; j < len(cous[i].Students); j++ {
			fmt.Println(cous[i].Students[j].Name)
		}
	}
	//fmt.Println(teas)
}

func Init() {
	dbURI := "host=localhost user=postgres dbname=nhan_local_database sslmode=disable password=Postgres port=5432"

	DbEntity, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect DB normal")
	}

	safe, _ := DbEntity.DB()

	defer safe.Close()

	DbEntity.AutoMigrate(&entity.Student{})
	DbEntity.AutoMigrate(&entity.Course{})
	DbEntity.AutoMigrate(&entity.Room{})
	DbEntity.AutoMigrate(&entity.Attendance{})
	DbEntity.AutoMigrate(&entity.StudentCourse{})
	DbEntity.AutoMigrate(&entity.Scheduler{})

	errJoin := DbEntity.SetupJoinTable(&entity.Student{}, "Courses", &entity.StudentCourse{})

	if errJoin != nil {
		panic(errJoin)
	}

	errJoin = DbEntity.SetupJoinTable(&entity.Room{}, "Courses", &entity.Scheduler{})

	if errJoin != nil {
		panic(errJoin)
	}

	fmt.Println("Migrate DB normal")
	//readDummy()
	//fmt.Println("Create Dummies")
	//createDummy()
}
