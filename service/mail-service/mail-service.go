package mail_service

import (
	"bytes"
	"github.com/go-gomail/gomail"
	"html/template"
)

var dialer *gomail.Dialer

type MailRequest struct {
	To      []string
	Subject string
	Body    string
}

func Init() error {
	dialer = gomail.NewDialer("smtp.yandex.com", 465, "no-reply@busmap.vn", "buSm@p.n0r3ply")

	_, err := dialer.Dial()
	return err
}

func (mr *MailRequest) SendEmail() error {
	m := gomail.NewMessage()

	m.SetHeader("From", "no-reply@busmap.vn")
	m.SetHeader("To", mr.To...)
	m.SetHeader("Subject", mr.Subject)
	m.SetBody("text/html", mr.Body)

	err := dialer.DialAndSend(m)
	return err
}

func (mr *MailRequest) ParseTemplate(templateFileName string, data interface{}) error {
	templateDirectory := "resources/templates/"
	tmpl, err := template.ParseFiles(templateDirectory + templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return err
	}

	mr.Body = buf.String()
	return nil
}
