package providers_test

import (
	"log"
	"os"
	"testing"

	"github.com/ivan-iver/hermes/models"
	p "github.com/ivan-iver/hermes/providers/mailchimp"
)

func TestMailchimpSendEmail(t *testing.T) {
	var provider = p.Mailchimp{}
	userEmail := os.Getenv("MANDRILL_USER")
	template := "<h1>notifik</h1><br><h3>Da click en el enlace o copialo en tu navegador<h3><h4>http://localhost:4000/users/sss/active</h4>"
	rec := []string{userEmail}
	email := models.Email{
		Subject:     "Bienvenido a Notifik",
		SenderEmail: "mauricio@optometrik.mx",
		SenderName:  "Alguien",
		Recipients:  rec,
		Content:     template,
	}
	log.Println(rec)
	if err := provider.SendEmail(email); err != nil {
		t.Error("Ya valio: ", err)
	}

	return

}
