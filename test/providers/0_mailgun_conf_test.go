package providers_test

import (
	"github.com/ivan-iver/hermes/models"
	"github.com/ivan-iver/hermes/providers/mailgun"
)

//ValidMailgunProvider
func ValidMailgunProvider() (e *mailgun.Mailgun) {
	provider := mailgun.Mailgun{}
	provider.Init()
	e = &provider
	return
}

//ValidEmail
func ValidMailgunEmail() (e *mailgun.Email) {
	email := &mailgun.Email{}
	sender := models.Sender{
		Email: "mailgun@hermes.com",
		Name:  "a friend",
	}
	email.AddSender(sender)
	email.AddContent(
		models.Content{
			Value: "Hollo from mailgun",
		},
	)
	email.AddSubject("welcome!!!")
	email.SetValues()
	email.AddRecipients(models.Recipients{To: []string{"mau.cdr.19@gmail.com", "mau_dsx2@hotmail.com"}})
	e = email
	return
}
