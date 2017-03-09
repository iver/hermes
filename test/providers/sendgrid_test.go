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
	template := "<h1>notifik</h1><br><h3>endgrid</h3>"
	rec := []string{userEmail}
	email := models.Email{
		Subject:     "Hola desde sendgrid",
		SenderEmail: "mau.cdr.19@gmail.com",
		SenderName:  "Alguien",
		Recipients:  rec,
		Content:     template,
	}

	if err := provider.SendEmail(email); err != nil {
		t.Error("Ya valio: ", err)
	}

	return

}
