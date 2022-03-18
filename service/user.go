package service

import (
	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(id string, req dto.UpdatePasswordRequest) error {
	user := repository.QueryUserBySID(id) // get ID from above
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return err
	}

	err = helper.ComparePassword(req.NewPass, req.ReNewPass)
	if err != nil {
		return err
	}

	newPassHash, err := helper.HashPassword(req.NewPass)

	if err != nil {
		return err
	}

	user.Password = newPassHash

	_, err = repository.Update(user)

	if err != nil {
		return err
	}

	return nil

}
