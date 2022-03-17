package repository

import (
	"github.com/pkg/errors"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryUserBySID(sid string) *entity.User {
	var user entity.User
	database.DbInstance.Where("id = ?", sid).First(&user)
	if user.ID == 0 {
		return nil
	}

	return &user
}

func Update(user *entity.User) (*entity.User, error) {

	err := database.DbInstance.Model(&user).Updates(&user).Error

	return user, errors.Wrap(err, "update user")
}

func UpdatePassword(passwordHash string, ID int64) error {

	err := database.DbInstance.Where("id= ?", ID).Updates(&entity.User{Password: passwordHash}).Error

	if err != nil {
		return errors.Wrap(err, "update user password")
	}

	return nil
}

func CheckEmailExist(email string) bool {
	user := &entity.User{}

	err := database.DbInstance.Where("email= ?", email).Find(&user).Error
	return err == nil
}
