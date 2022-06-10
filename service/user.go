package service

import (
	"errors"
	"math/rand"

	"github.com/smartschool/helper"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
	mail_service "github.com/smartschool/service/mail-service"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(id string, req dto.UpdatePasswordRequest) error {
	user := repository.QueryUserBySID(id) // get ID from above
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return err
	}

	err = helper.CompareOldNewPass(req.Password, req.NewPass)
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

func ChangePasswordFirstTime(id string, req dto.ChangePasswordFirstTimeRequest) error {
	user := repository.QueryUserBySID(id)

	if user.IsActivate {
		return errors.New("user is not allowed to change password")
	}

	err := helper.ComparePassword(req.NewPass, req.ReNewPass)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.NewPass))
	if err == nil {
		return errors.New("can not update on the same password")
	}

	newPassHash, err := helper.HashPassword(req.NewPass)
	if err != nil {
		return err
	}

	user.Password = newPassHash
	user.IsActivate = true

	_, err = repository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func ResetPassword(req dto.ResetPasswordRequest) error {
	user := repository.QueryUserByEmail(req.Email)
	if user == nil {
		return errors.New("User not found")
	}

	newPassword := generatePassword(10)
	newPasswordHash, err := helper.HashPassword(newPassword)
	if err != nil {
		return errors.New("Error hashing password")
	}

	user.Password = newPasswordHash
	_, err = repository.Update(user)
	if err != nil {
		return errors.New("Error updating password")
	}

	mr := &mail_service.MailRequest{
		To:      []string{user.Email},
		Subject: "Your password for Student Connect has been reset",
	}
	resetPasswordEmailData := &dto.ResetPasswordEmailData{
		NewPassword: newPassword,
	}
	err = mr.ParseTemplate("mail-reset-password.txt", resetPasswordEmailData)
	if err != nil {
		return errors.New("Error generating email")
	}

	err = mr.SendEmail()
	if err != nil {
		return errors.New("Error sending email")
	}

	return nil
}

func generatePassword(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
