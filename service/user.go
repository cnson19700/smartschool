package service

import (
	"errors"

	"github.com/smartschool/helper"
	"github.com/smartschool/repository"
)

type UpdatePasswordRequest struct {
	Password  string `json:"password"`
	NewPass   string `json:"new_password"`
	ReNewPass string `json:"re_new_password"`
}

func UpdatePassword(id string, req UpdatePasswordRequest) error {
	if req.NewPass != req.ReNewPass {
		return errors.New("2 password not matches")
	}

	isPass, newPass := helper.CheckFormatValue("password", req.NewPass)
	if !isPass {
		return errors.New("Wrong password format")
	}
	if len(newPass) < 8 {
		return errors.New("password must have 8 characters")
	}

	isPass, oldPass := helper.CheckFormatValue("password", req.Password)
	if !isPass {
		return errors.New("Wrong password format")
	}

	user := repository.QueryUserBySID(id) // get ID from above
	isPassTrue := helper.VerifyPassword(oldPass, user.Password)
	if !isPassTrue {
		return errors.New("Wrong password!!")
	}

	newPassHash, err := helper.HashPassword(newPass)

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
