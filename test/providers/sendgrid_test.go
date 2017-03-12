package providers_test

import (
	"os"
	"testing"
	p "github.com/ivan-iver/hermes/providers/sendgrid"
)

func TestMailchimpSendEmail(t *testing.T) {
	var provider = p.Sendgrid{}
	userEmail := os.Getenv("SENDGRID_USER")
	senderName:= "Un amigo"
	senderEmail:= "sendgrid@hermes.mx"
	subject:= "Un saludo"
	template := "<div><h1>Hola desde sendgrid<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{userEmail, "mau16@ciencias.unam.mx"}
    provider.Init()
    email,_:= provider.NewEmail(senderEmail,senderName,subject,template)
	email.AddRecipients(recipients...)
	if err := provider.SendEmail(email); err != nil {
		t.Error("hermes: error in provider", err)
	}

	return

}
