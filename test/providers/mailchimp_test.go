package providers_test

import (
	"os"
	"testing"

	"github.com/ivan-iver/hermes/models"
	p "github.com/ivan-iver/hermes/providers/mailchimp"
)

func TestMailchimpSendEmail(t *testing.T) {
	var provider = p.Mailchimp{}
	userEmail := os.Getenv("MANDRILL_USER")
	template := "<div><h1>Hola desde mailchimp<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{userEmail, "mau.cdr.19@gmail.com"}
	email := models.Email{
		Subject:     "Hola desde mailchimp",
		SenderEmail: "mauricio@optometrik.mx",
		SenderName:  "Alguien",
		Recipients:  recipients,
		Template:    template,
	}

	if err := provider.SendEmail(email); err != nil {
		t.Error("hermes: error in provider", err)
	}

	return

}
