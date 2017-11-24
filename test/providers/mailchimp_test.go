package providers_test

import (
	"github.com/iver/hermes/models"
	"github.com/iver/hermes/providers/mailchimp"
	"log"
	"testing"
)

func TestCreateMailchimpProviderOK(t *testing.T) {
	var provider = mailchimp.Mailchimp{}
	if err := provider.Init(); err != nil {
		t.Error("provider:Init()-", err)
	}

	if name := provider.GetName(); name != "mailchimp" {
		log.Println("Provider name:", name)
		t.Error("provider:Wrong Name")
	}
	return
}

func TestCreateMailchimpProviderFail(t *testing.T) {
	mailchimp.Cfgfile = "other.conf"
	defer func() {
		if r := recover(); r != nil {
			mailchimp.Cfgfile = `provider.conf`
		}
		return
	}()
	var provider = mailchimp.Mailchimp{}
	if err := provider.Init(); err == nil {
		t.Error("provider:Init()-", err)
	}
	return
}

func TestCreateMailchimpEmailOK(t *testing.T) {
	var provider = ValidMailchimpProvider()
	var sender = models.Sender{Name: "Un Friend", Email: "mailchimp@hermes.mx"}
	var subject = "Welcome!!"
	var content = models.Content{Value: "Hello from mailchimp"}
	if _, err := provider.NewEmail(sender, subject, content); err != nil {
		t.Error("provider:NewEmail()", err)
	}
	return
}

func TestCreateMailchimpEmailFail(t *testing.T) {
	var provider = ValidMailchimpProvider()
	var sender = "sender"
	var subject = "Welcome!!"
	var content = models.Content{Value: "Hello from mailchimp"}
	if _, err := provider.NewEmail(sender, subject, content); err == nil {
		t.Error("provider:NewEmail()", err)
	}
	return
}

func TestMailchimpSendEmailOK(t *testing.T) {
	var provider = ValidMailchimpProvider()
	var emailM = ValidMailchimpEmail()
	if err := provider.SendEmail(emailM); err != nil {
		t.Error("provider:SendEmail()-", err)
	}
	return
}

func TestMailchimpSendEmailFail(t *testing.T) {
	//TODO:
	return
}
