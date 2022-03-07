package api_students

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/smartschool/database"
	"github.com/smartschool/entity"
	"go.elastic.co/apm/model"
)

func GetStudentByID(c *gin.Context, id string) (*entity.Student, error) {
	stu := &entity.Student{}
	err := database.DbInstance.Where("student_id = ?", id).First(&stu).Error
	if err != nil {
		return nil, errors.Wrap(err, "get user by id")
	}
	return stu, nil
}

func GetAll(c *gin.Context) ([]entity.Student, error) {
	listStu := []entity.Student{}
	database.DbInstance.Find(&listStu)
	return listStu, nil
}

func GetEmail(c *gin.Context, email string) (*entity.User, error) {
	user := &entity.User{}

	err := database.DbInstance.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return nil, errors.Wrap(err, "get user by email")
	}

	return user, nil
}

func CheckEmailExist(c *gin.Context, email string) bool {
	user := &model.User{}

	err := database.DbInstance.Where("email= ?", email).Find(&user).Error

	if err != nil {
		return false
	}

	return true
}
