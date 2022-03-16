package excel

import (
	"github.com/gin-gonic/gin"
	"github.com/smartschool/database"
	"github.com/smartschool/helper"
	"github.com/smartschool/model/entity"
)

const USER_SHEET_NAME = "User"

func ImportUser(c *gin.Context) {
	w := c.Writer

	excel, err := PreprocessImport(c)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	rows := excel.GetRows(USER_SHEET_NAME)
	users := make([]entity.User, 0)
	for i := 1; i < len(rows); i += 1 {
		row := rows[i]

		var user entity.User

		user.Email = row[0]

		hashedPassword, err := helper.HashPassword(row[1])
		if err != nil {
			w.Write([]byte("Error hashing password"))
			return
		}
		user.Password = hashedPassword

		users = append(users, user)
	}

	database.DbInstance.Create(&users)
}
