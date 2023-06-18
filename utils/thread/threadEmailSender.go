package thread

import (
	"authorizationGRPC/internal/config"
	"authorizationGRPC/utils/emailSender"
)

type EmailSenderThread struct {
	thread      *Thread
	emailSender *emailSender.EmailSender
}

func NewThreadEmailSender(c config.Config) *EmailSenderThread {
	emailSender := emailSender.NewEmailSender(c.Mail.Identity, c.Mail.Login,
		c.Mail.From, c.Mail.Password, c.Mail.SmtpHost, c.Mail.SmtpPort)
	thread := New(emailSender.Send)
	return &EmailSenderThread{
		thread:      thread,
		emailSender: emailSender,
	}
}
func (e *EmailSenderThread) SetMessage(m emailSender.Email) {
	e.thread.Push(m)
}
