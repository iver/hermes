package providers_test

import (
	"os"
	"testing"
	p "github.com/ivan-iver/hermes/providers/mailchimp"
)

func TestMailchimpSendEmail(t *testing.T) {
	var err error
	var provider = p.Mailchimp{}
	userEmail := os.Getenv("MANDRILL_USER")
	senderName:= "Un amigo"
	senderEmail:= "mailchimp@hermes.mx"
	subject:= "Un saludo"
	content := "<div><h1>Hola desde mailchimp<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{userEmail, "mau.cdr.19@gmail.com"}
    if err=provider.Init(); err != nil{
       t.Error("provider:Init()", err)
	}
    email,_:= provider.NewEmail(senderEmail,senderName,subject,content)
	email.AddRecipients(recipients...)
	if err := provider.SendEmail(email); err != nil {
	   t.Error("provider:SendEmail()-", err)
	}
	return

}
