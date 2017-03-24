package providers_test

import (
	"testing"
	"github.com/ivan-iver/hermes/models"
	"github.com/ivan-iver/hermes/providers/mailgun"
)

func TestMailgunSendEmail(t *testing.T) {
	var err error
	var provider = mailgun.Mailgun{}
	var emailM *mailgun.Email
	var sender = models.Sender{Name:"Un Amigo",Email:"mailgun@hermes.mx"}
	subject := "Un saludo"
    var content = models.Content{Value:"Hola desde mailgun"}
	var recipients = models.Recipients{To:[]string{"mau.cdr.19@gmail.com", "mau16@ciencias.unam.mx"}};
    if err = provider.Init(); err !=nil{
		 t.Error("provider:Init()-", err)
	}
    email,err:= provider.NewEmail(sender,subject,content)
	emailM=email.(*mailgun.Email)
	emailM.AddRecipients(recipients)
	if err = provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()-", err)
	}
	return

}

