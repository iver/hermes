package providers_test

import (
	//"os"
	"testing"
	p "github.com/ivan-iver/hermes/providers/mailgun"
)

func TestMailchimpSendEmail(t *testing.T) {
	var provider = p.Mailgun{}
	/*userEmail := os.Getenv("MG_EMAIL_TO")*/
	senderName:= "Un amigo"
	senderEmail:= "mailgun@hermes.mx"
	subject:= "Un saludo"
	template := "<div><h1>Hola desde mailgun<h1><h4>Template desde hermes</h4></div>"
	recipients := []string{"mau.cdr.19@gmail.com"}
    provider.Init()
    email,_:= provider.NewEmail(senderEmail,senderName,subject,template)
	email.AddRecipients(recipients...)
	if err := provider.SendEmail(email); err != nil {
		t.Error("hermes: error in provider", err)
	}

	return

}

