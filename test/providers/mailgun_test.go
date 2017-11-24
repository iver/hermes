package providers_test

import (
	"github.com/iver/hermes/models"
	"github.com/iver/hermes/providers/mailgun"
	"log"
	"testing"
)

func TestCreateMailgunProviderOK(t *testing.T) {
	var provider = mailgun.Mailgun{}
	if err := provider.Init(); err != nil {
		t.Error("provider:Init()-", err)
	}

	if name := provider.GetName(); name != "mailgun" {
		log.Println("Provider name:", name)
		t.Error("provider:Wrong Name")
	}
	return
}

func TestCreateMailgunProviderFail(t *testing.T) {
	mailgun.Cfgfile = "other.conf"
	defer func() {
		if r := recover(); r != nil {
			mailgun.Cfgfile = `provider.conf`
		}
		return
	}()
	var provider = mailgun.Mailgun{}
	if err := provider.Init(); err == nil {
		t.Error("provider:Init()-", err)
	}
	return
}

func TestCreateMailgunEmailOK(t *testing.T) {
	var provider = ValidMailgunProvider()
	var sender = models.Sender{Name: "Un Friend", Email: "mailgun@hermes.mx"}
	var subject = "Welcome!!"
	var content = models.Content{Value: "Hello from mailgun"}
	if _, err := provider.NewEmail(sender, subject, content); err != nil {
		t.Error("provider:NewEmail()", err)
	}
	return
}

func TestCreateMailgunEmailFail(t *testing.T) {
	var provider = ValidMailgunProvider()
	var sender = "sender"
	var subject = "Welcome!!"
	var content = models.Content{Value: "Hello from mailgun"}
	if _, err := provider.NewEmail(sender, subject, content); err == nil {
		t.Error("provider:NewEmail()", err)
	}
	return
}

func TestMailgunSendEmailOK(t *testing.T) {
	var provider = ValidMailgunProvider()
	var emailM = ValidMailgunEmail()
	if err := provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()-", err)
	}
	return
}

/*func TestMailgunSendEmailFail(t *testing.T) {
	//TODO:
	return
}
*/
