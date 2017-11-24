package hermes_test

import (
	"github.com/iver/hermes/models"
)

type Email struct {
	Sender     models.Sender
	Subject    string
	Content    models.Content
	Recipients models.Recipients
}

//CorrectEmail return valid information to an email
func CorrectEmail() (e *Email) {
	e = &Email{
		Sender:     models.Sender{Name: "A friend", Email: "hermes@hermes.com"},
		Content:    models.Content{Value: "Hello.......!"},
		Subject:    "Test with hermes",
		Recipients: models.Recipients{To: []string{"mau.cdr.19@gmail.com", "mau_dsx2@hotmail.com"}},
	}
	return
}
