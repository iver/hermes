package providers_test

import (
	"github.com/iver/hermes/models"
	"github.com/iver/hermes/providers/mailchimp"
)

//ValidMailchimpProvider
func ValidMailchimpProvider() (e *mailchimp.Mailchimp) {
	provider := mailchimp.Mailchimp{}
	provider.Init()
	e = &provider
	return
}

//ValidEmail
func ValidMailchimpEmail() (e *mailchimp.Email) {
	email := &mailchimp.Email{}
	email.AddSender(
		models.Sender{
			Email: "mailchimp@hermes.com",
			Name:  "a friend",
		},
	)
	email.AddContent(
		models.Content{
			Value: "Hollo from mailchimp",
		},
	)
	email.AddSubject("welcome!!!")
	email.AddRecipients(models.Recipients{To: []string{"mau.cdr.19@gmail.com", "mau_dsx2@hotmail.com"}})
	e = email
	return
}
