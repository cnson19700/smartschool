package service

import (
	"errors"

	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
)

func UpdatePassword(id string, req dto.UpdatePasswordRequest) error {
	err := helper.ComparePassword(req.Password, req.NewPass)
	if err != nil {
		return errors.New("Fail to compare password")
	}

	user := repository.QueryUserBySID(id) // get ID from above
	isPassTrue := helper.VerifyPassword(req.Password, user.Password)
	if !isPassTrue {
		return errors.New("Wrong password!!")
	}

	newPassHash, err := helper.HashPassword(req.NewPass)

	if err != nil {
		return errors.New("Password hash failed")
	}

	user.Password = newPassHash

	_, err = repository.Update(user)
	if err != nil {
		return errors.New("Fail to update password")
	}

	return nil

}
