package providers_test

import (
	"os"
	"testing"

	"github.com/ivan-iver/hermes/models"
	p "github.com/ivan-iver/hermes/providers/mailgun"
)

func TestMailgunSendEmail(t *testing.T) {
	var provider = p.Mailgun{}
	userEmail := os.Getenv("MG_EMAIL_TO")
    template := "<div><h1>Hola desde Mailgun<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{userEmail,"mau.cdr.19@gmail.com"}
	email := models.Email{
		Subject:     "Hola desde mailgun",
		SenderEmail: "person@hermes.mx",
		SenderName:  "Alguien",
		Recipients:  recipients,
		Content:     template,
	}

	if err := provider.SendEmail(email); err != nil {
		t.Error("hermes: error in provider", err)
	}

	return

}
