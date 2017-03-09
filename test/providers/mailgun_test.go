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
	template := "<h1>notifik</h1><br><h4>Hola desde mailgun</h4>"
	rec := []string{userEmail}
	email := models.Email{
		Subject:     "Hola desde mailgun",
		SenderEmail: "person@hermes.mx",
		SenderName:  "Alguien",
		Recipients:  rec,
		Content:     template,
	}

	if err := provider.SendEmail(email); err != nil {
		t.Error("Ya valio: ", err)
	}

	return

}
