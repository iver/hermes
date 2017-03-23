package providers_test

import (
	"testing"
	"github.com/ivan-iver/hermes/models"
	"github.com/ivan-iver/hermes/providers/sendgrid"
)

func TestSendgridSendEmail(t *testing.T) {
	var err error
	var provider = sendgrid.Sendgrid{}
	var emailM *sendgrid.Email
	var sender = models.Sender{Name:"Un Amigo",Email:"sendgrid@hermes.mx"}
	subject := "Un saludo"
    var content = models.Content{Value:"Hola desde sendgrid"}
	var recipients = models.Recipients{To:[]string{"mau.cdr.19@gmail.com", "mau16@ciencias.unam.mx"}};
    if err = provider.Init(); err !=nil{
		 t.Error("provider:Init()-", err)
	}
    email,err:= provider.NewEmail(sender,subject,content)
	emailM=email.(*sendgrid.Email)
	emailM.AddRecipients(recipients)
	if err = provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()-", err)
	}

	return

}
