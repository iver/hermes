package providers_test

import (
	"testing"
	"github.com/ivan-iver/hermes/providers/sendgrid"
)

func TestMailchimpSendEmail(t *testing.T) {
	var err error
	var provider = sendgrid.Sendgrid{}
	var emailM *sendgrid.Email
	senderName:= "Un amigo"
	senderEmail:= "sendgrid@hermes.mx"
	subject:= "Un saludo"
	template := "<div><h1>Hola desde sendgrid<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{"mau.cdr.19@gmail.com", "mau16@ciencias.unam.mx"}
    if err = provider.Init(); err !=nil{
		 t.Error("provider:Init()-", err)
	}
    email,err:= provider.NewEmail(senderEmail,senderName,subject,template)
	emailM=email.(*sendgrid.Email)
	emailM.AddRecipients(recipients...)
	if err = provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()-", err)
	}

	return

}
