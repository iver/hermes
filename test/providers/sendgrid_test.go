package providers_test

import (
	"os"
	"testing"

	"github.com/ivan-iver/hermes/models"
	p "github.com/ivan-iver/hermes/providers/sendgrid"
)

func TestSendGridSendEmail(t *testing.T) {
	var provider = p.Sendgrid{}
	userEmail := os.Getenv("SENDGRID_USER")
	template := "<div><h1>Hola desde sendgrid<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{userEmail, "mau16@ciencias.unam.mx"}

	email := models.Email{
		Subject:     "Hola desde sendgrid",
		SenderEmail: "person@hermes.mx",
		SenderName:  "Alguien",
		Recipients:  recipients,
		Template:    template,
	}
   
	if err := provider.SendEmail(email); err != nil {
		t.Error("hermes: error in provider", err)
	}
   
	return

}
