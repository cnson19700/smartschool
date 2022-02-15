package database

import (
	"fmt"

	"github.com/smartschool/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbEntity *gorm.DB

func Init() {
	ConnectDatabase()
	MigrateDatabase()
	Close()
}

func Close() {
	safe, _ := DbEntity.DB()

	defer safe.Close()
}

func ConnectDatabase() {
	dbURI := "host=13.228.244.196 port=5432 user=busmapdb dbname=phenikaamaas_attendancedb sslmode=disable password=frjsdfhaflpzlcdzgnfvuxkdwiiiiklpojzowxajmendeeoqtbzyrgi"

	var err error
	DbEntity, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect DB normal")
	}
}

func MigrateDatabase() {
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
}
