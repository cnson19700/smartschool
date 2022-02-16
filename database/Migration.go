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
	Close()
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
	DbInstance.AutoMigrate(&entity.Attendance{})
	DbInstance.AutoMigrate(&entity.StudentCourse{})
	DbInstance.AutoMigrate(&entity.Scheduler{})
	DbInstance.AutoMigrate(&entity.Device{})

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
