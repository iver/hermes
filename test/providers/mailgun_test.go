package providers_test

import (
	//"os"
	"testing"
	p "github.com/ivan-iver/hermes/providers/mailgun"
)

func TestMailchimpSendEmail(t *testing.T) {
	var provider = p.Mailgun{}
	var err error
	senderName:= "Un amigo"
	senderEmail:= "mailgun@hermes.mx"
	subject:= "Un saludo"
	template := "<div><h1>Hola desde mailgun<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{"mau.cdr.19@gmail.com"}
    if err = provider.Init(); err != nil{
	   t.Error("provider:Init()", err)	
	}
    email,_:= provider.NewEmail(senderEmail,senderName,subject,template)
	email.AddRecipients(recipients...)
	if err := provider.SendEmail(email); err != nil {
		t.Error("provider:SendEmail()", err)
	}

	return

}

