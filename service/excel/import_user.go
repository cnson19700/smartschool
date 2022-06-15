package excel

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/apptypes"
	"github.com/smartschool/database"
	"github.com/smartschool/helper"
	"github.com/smartschool/model/entity"
)

const USER_SHEET_NAME = "User"

// func ImportUser(c *gin.Context) {
// 	w := c.Writer

// 	excel, err := PreprocessImport(c)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 	}

// 	rows := excel.GetRows(USER_SHEET_NAME)
// 	users := make([]entity.User, 0)
// 	for i := 1; i < len(rows); i += 1 {
// 		row := rows[i]

// 		var user entity.User

// 		user.Email = row[0]

// 		hashedPassword, err := helper.HashPassword(row[1])
// 		if err != nil {
// 			w.Write([]byte("Error hashing password"))
// 			return
// 		}
// 		user.Password = hashedPassword

// 		users = append(users, user)
// 	}

// 	database.DbInstance.Create(&users)
// }

func ImportUser(c *gin.Context) {
	//w := c.Writer

	excel, err := PreprocessImport(c, "public/user_import/")
	if err != nil {
		//w.Write([]byte(err.Error()))
		return
	}

	//users := make([]entity.User, 0)
	var day, month, year string
	for idr, row := range excel.GetRows(USER_SHEET_NAME) {

		if idr == 0 {
			continue
		}

		serialTime, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			continue
		}
		DoB := ExcelSerialDateToTime(serialTime)

		if DoB.Day() < 10 {
			day = "0" + strconv.Itoa(DoB.Day())
		} else {
			day = strconv.Itoa(DoB.Day())
		}

		if int(DoB.Month()) < 10 {
			month = "0" + strconv.Itoa(int(DoB.Month()))
		} else {
			month = strconv.Itoa(int(DoB.Month()))
		}

		year = strconv.Itoa(DoB.Year())

		hashedPassword, bcryptError := helper.HashPassword(day + month + year)
		if bcryptError != nil {
			continue
		}

		var user entity.User
		var dummy entity.User

		user = entity.User{
			Username:    row[1],
			Password:    hashedPassword,
			FirstName:   row[2],
			DateOfBirth: DoB,
			Email:       row[8],
			PhoneNumber: row[15],
			RoleID:      apptypes.StudentRole,
			FacultyID:   1,
			IsActivate:  false,
			Student: &entity.Student{
				StudentID: row[1],
				Batch:     row[6],
			},
		}

		res := database.DbInstance.Where("user_name = ?", user.Username).Find(&dummy)
		if res.Error != nil {
			continue
		}
		if res.RowsAffected == 0 {
			database.DbInstance.Create(&user)
		} else {
			dummy.Password = user.Password
			dummy.FirstName = user.FirstName
			dummy.DateOfBirth = user.DateOfBirth
			dummy.Email = user.Email
			dummy.PhoneNumber = user.PhoneNumber
			dummy.RoleID = user.RoleID
			dummy.FacultyID = user.FacultyID
			dummy.Student = user.Student

			database.DbInstance.Save(&dummy)
		}
	}
}

func ExcelSerialDateToTime(serial float64) time.Time {
	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	convertedTime := excelEpoch.Add(time.Duration(serial * float64(24*time.Hour)))

	return convertedTime
}
