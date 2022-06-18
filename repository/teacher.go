package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryTeacherByID(id string) (*entity.Teacher, error) {
	teacher := &entity.Teacher{}
	err := database.DbInstance.Where("id = ?", id).First(teacher).Error
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func QueryTeacherByName(name string) (*entity.User, bool, error) {
	var user entity.User

	result := database.DbInstance.Where("first_name = ?", name).
		Limit(1).Find(&user)

	return &user, result.RowsAffected == 0, result.Error
}

func GetIDFromTeacherID(teacher_id string) (uint, error){
	teacher := &entity.Teacher{}
	err := database.DbInstance.Where("teacher_id = ?", teacher_id).First(teacher).Error
	if err != nil {
		return 0, err
	}

	return teacher.ID, nil
}