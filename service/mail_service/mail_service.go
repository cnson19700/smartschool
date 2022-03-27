package mailservice

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/smartschool/model/dto"

	gomail "gopkg.in/mail.v2"
)

var dialer *gomail.Dialer

func NewRequest(to string, subject, body string) *dto.RequestMail {
	return &dto.RequestMail{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func SendEmail(r *dto.RequestMail) (bool, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "personal-email") //example
	m.SetHeader("To", r.To)
	m.SetHeader("Subject", r.Subject)
	m.SetBody("text/html", r.Body)

	// Send the email to Bob, Cora and Dan.
	if err := dialer.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}

func ParseTemplate(r *dto.RequestMail, templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}

func Initialize() {
	dialer = gomail.NewDialer("smtp.gmail.com", 587, "from", "password")

	_, err := dialer.Dial()
	if err == nil {
		fmt.Println("Connected to Mail Service!")
	} else {
		fmt.Println("Mail Service connection errors!", err)
	}
}
