package providers_test

import (
	"log"
	"testing"

	"github.com/ivan-iver/hermes/models"
	"github.com/ivan-iver/hermes/providers/sendgrid"
)

func TestCreateSendgridProviderOK(t *testing.T) {
	var provider = sendgrid.Sendgrid{}
	if err := provider.Init(); err != nil {
		t.Error("provider:Init()-", err)
	}

	if name := provider.GetName(); name != "sendgrid" {
		log.Println("Provider name:", name)
		t.Error("provider:Wrong Name")
	}
	return
}

func TestCreateSendgridProviderFail(t *testing.T) {
	sendgrid.Cfgfile = "other.conf"
	defer func() {
		if r := recover(); r != nil {
			sendgrid.Cfgfile = `provider.conf`
		}
		return
	}()
	var provider = sendgrid.Sendgrid{}
	if err := provider.Init(); err == nil {
		t.Error("provider:Init()-", err)
	}
	return
}

func TestCreateSendgridEmailOK(t *testing.T) {
	var provider = ValidSendgridProvider()
	var sender = models.Sender{Name: "Un Friend", Email: "sendgrid@hermes.mx"}
	var subject = "Welcome!!"
	var content = models.Content{Value: "Hello from sendgrid"}
	if _, err := provider.NewEmail(sender, subject, content); err != nil {
		t.Error("provider:NewEmail()", err)
	}
	return
}

func TestCreateSendgridEmailFail(t *testing.T) {
	var provider = ValidSendgridProvider()
	var sender = "sender"
	var subject = "Welcome!!"
	var content = models.Content{Value: "Hello from sendgrid"}
	if _, err := provider.NewEmail(sender, subject, content); err == nil {
		t.Error("provider:NewEmail()", err)
	}
	return
}

func TestSendgridSendEmailOK(t *testing.T) {
	var provider = ValidSendgridProvider()
	var emailM = ValidSendgridEmail()
	if err := provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()-", err)
	}
	return
}

func TestSendgridSendEmailFail(t *testing.T) {
	//TODO:
	return
}
