package excel

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

	excel, err := PreprocessImport(c, "public/user_import/")
	if err != nil {
		return
	}

	headReader := false
	idc := 0
	var day, month, year string
	for _, row := range excel.GetRows(USER_SHEET_NAME) {

		if !headReader {
			for local_idc, col := range row {
				if strings.ToLower(col) == apptypes.ImportUser_Marker {
					headReader = true
					idc = local_idc
					break
				}
			}
			continue
		}

		var DoB time.Time
		dateType := reflect.TypeOf(row[idc+4]).Kind()
		if dateType == reflect.String {
			DoB, err = helper.StringToDateUTC(row[idc+4])
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			serialTime, err := strconv.ParseFloat(row[idc+4], 64)
			if err != nil {
				continue
			}
			DoB = ExcelSerialDateToTime(serialTime)
		}

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

		studentGender := apptypes.DB_Gender_Female
		if strings.ToLower(row[idc+5]) == apptypes.ImportUser_Gender_Male {
			studentGender = apptypes.DB_Gender_Male
		}

		var user entity.User
		var dummyUser entity.User

		user = entity.User{
			Username:    row[idc],
			Password:    hashedPassword,
			LastName:    row[idc+1],
			FirstName:   row[idc+2],
			DateOfBirth: DoB,
			Email:       row[idc+7],
			RoleID:      apptypes.StudentRole,
			Gender:      studentGender,
			FacultyID:   1,
			IsActivate:  false,
			Student: &entity.Student{
				StudentID: row[idc],
				Batch:     row[idc+6],
			},
		}

		res := database.DbInstance.Where("user_name = ?", user.Username).Find(&dummyUser)
		if res.Error != nil {
			continue
		}
		if res.RowsAffected == 0 {
			database.DbInstance.Create(&user)
		} else {
			dummyUser.Password = user.Password
			dummyUser.LastName = user.LastName
			dummyUser.FirstName = user.FirstName
			dummyUser.DateOfBirth = user.DateOfBirth
			dummyUser.Gender = user.Gender
			dummyUser.Email = user.Email
			dummyUser.RoleID = user.RoleID
			dummyUser.FacultyID = user.FacultyID
			dummyUser.Student = user.Student

			database.DbInstance.Save(&dummyUser)
		}
	}
}

func ExcelSerialDateToTime(serial float64) time.Time {
	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	convertedTime := excelEpoch.Add(time.Duration(serial * float64(24*time.Hour)))

	return convertedTime
}
