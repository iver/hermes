package providers_test

import (
	"github.com/iver/hermes/models"
	"github.com/iver/hermes/providers/sendgrid"
)

//ValidSendgridProvider
func ValidSendgridProvider() (e *sendgrid.Sendgrid) {
	provider := sendgrid.Sendgrid{}
	provider.Init()
	e = &provider
	return
}

//ValidEmail
func ValidSendgridEmail() (e *sendgrid.Email) {
	emailp := sendgrid.NewEmail()
	email := &emailp
	sender := models.Sender{
		Email: "sendgrid@hermes.com",
		Name:  "a friend",
	}
	email.AddSender(sender)
	email.AddContent(
		models.Content{
			Value: "Hollo from sendgrid",
		},
	)
	email.AddSubject("welcome!!!")
	email.AddRecipients(models.Recipients{To: []string{"mau.cdr.19@gmail.com", "mau_dsx2@hotmail.com"}})
	e = email
	return
}
