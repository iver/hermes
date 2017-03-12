package providers_test

import (
	"os"
	"testing"
	p "github.com/ivan-iver/hermes/providers/mailchimp"
)

func TestMailchimpSendEmail(t *testing.T) {
	var provider = p.Mailchimp{}
	userEmail := os.Getenv("MANDRILL_USER")
	senderName:= "Un amigo"
	senderEmail:= "mailchimp@hermes.mx"
	subject:= "Un saludo"
	content := "<div><h1>Hola desde mailchimp<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{userEmail, "mau.cdr.19@gmail.com"}
    provider.Init()
    email,_:= provider.NewEmail(senderEmail,senderName,subject,content)
	email.AddRecipients(recipients...)
	if err := provider.SendEmail(email); err != nil {
		t.Error("hermes: error in provider", err)
	}

	return

}
