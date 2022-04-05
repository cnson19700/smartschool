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
