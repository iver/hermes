package providers_test

import (
	"testing"
	"github.com/ivan-iver/hermes/providers/mailgun"
)

func TestMailgunSendEmail(t *testing.T) {
	var err error
	var provider = mailgun.Mailgun{}
	var emailM *mailgun.Email
	senderName:= "Un amigo"
	senderEmail:= "mailgun@hermes.mx"
	subject:= "Un saludo"
	template := "<div><h1>Hola desde mailgun<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{"mau.cdr.19@gmail.com"}
    if err = provider.Init(); err != nil{
	   t.Error("provider:Init()", err)	
	}
    email,err := provider.NewEmail(senderEmail,senderName,subject,template)
	emailM = email.(*mailgun.Email)
	emailM.AddRecipients(recipients...)
	if err := provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()", err)
	}

	return

}

