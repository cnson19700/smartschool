package repository

import (
	"github.com/pkg/errors"
	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
)

func QueryUserBySID(sid string) *entity.User {
	var user entity.User
	err := database.DbInstance.Where("id = ?", sid).First(&user).Error
	if err != nil {
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

func QueryUserRoleIDByUserID(id uint) (uint, error) {
	var role_id uint
	result := database.DbInstance.Table("users").Select("role_id").Where("id = ?", id).Find(&role_id)

	return role_id, result.Error
}

func QueryUserRoleDetailByRoleID(id uint) (string, error) {
	var role_title string
	result := database.DbInstance.Table("roles").Select("title").Where("id = ?", id).Find(&role_title)

	return role_title, result.Error
}

func QueryUserByEmail(email string) *entity.User {
	var user entity.User
	err := database.DbInstance.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil
	}

	return &user
}

func QueryListUserNameInfo(ids []uint) ([]dto.UserFullNameInfo, bool, error) {
	var users []dto.UserFullNameInfo
	result := database.DbInstance.Table("users").Where("id IN ? AND deleted_at IS NULL", ids).Find(&users)

	return users, result.RowsAffected == 0, result.Error
}

func QueryUserNameInfo(id uint) (dto.UserFullNameInfo, error) {
	var user dto.UserFullNameInfo
	result := database.DbInstance.Table("users").Where("id = ? AND deleted_at IS NULL", id).First(&user)

	return user, result.Error
}
