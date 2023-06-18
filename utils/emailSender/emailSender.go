package emailSender

import (
	"net/smtp"
)

type EmailSender struct {
	identity string
	from     string
	password string
	smtpHost string
	smtpPort string
	auth     smtp.Auth
}

type Email struct {
	Message string
	To      []string
}

func NewEmailSender(identity, login, from, password, host, port string) *EmailSender {
	auth := smtp.PlainAuth(identity, login, password, host)
	return &EmailSender{
		identity: identity,
		from:     from,
		password: password,
		smtpHost: host,
		smtpPort: port,
		auth:     auth,
	}
}

func (e *EmailSender) Send(eml any) error {
	em := eml.(Email)
	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, e.auth, e.from, em.To, []byte(em.Message))
	if err != nil {
		return err
	}
	return nil
}
