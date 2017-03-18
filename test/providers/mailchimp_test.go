package providers_test

import (
	"testing"
    "github.com/ivan-iver/hermes/providers/mailchimp"
)

func TestMailchimpSendEmail(t *testing.T) {
	var err error
	var provider = mailchimp.Mailchimp{} 
	var emailM *mailchimp.Email
	senderName:= "Un amigo"
	senderEmail:= "mailchimp@hermes.mx"
	subject:= "Un saludo"
	content := "<div><h1>Hola desde mailchimp<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{"mau_dsx2@hotmail.com", "mau.cdr.19@gmail.com"}
    if err=provider.Init(); err != nil{
       t.Error("provider:Init()", err)
	}
    email,err := provider.NewEmail(senderEmail,senderName,subject,content)
    emailM = email.(*mailchimp.Email)
	emailM.AddRecipients(recipients...)
	if err := provider.SendEmail(emailM); err != nil {
	   t.Error("provider:SendEmail()-", err)
	}
	return

}
