package dto

type MailType int

// List of Mail Types we are going to send.
const (
	MailConfirmation MailType = iota + 1
	PassReset
)

// MailData represents the data to be sent to the template of the mail.
type MailData struct {
	Username string `json:"user_name"`
	Code     string `json:"code"`
}

// Mail represents a email request
type RequestMail struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	mailType MailType
	data     *MailData
}
