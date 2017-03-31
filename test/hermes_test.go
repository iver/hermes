package hermes_test

import (
	"log"
	"testing"

	"github.com/ivan-iver/hermes"
)

var (
	emailProvider *hermes.EmailProvider
	configFile = "provider.conf"
)

func TestCreateHermesOK(t *testing.T) {
	var err error
	if emailProvider, err = hermes.New(configFile); err != nil {
		t.Error("hermes:New()", err)
	}
	if name := emailProvider.Selected(); name != "mailchimp" {
		log.Println("ProviderSelected:", name)
		t.Error("Wrong provider")
	}
	return
}

func TestCreateHermesFail(t *testing.T) {
	//TODO..
	return
}

func TestSendEmailOK(t *testing.T) {
	var err error
	iEmail := CorrectEmail()
	email, err := emailProvider.NewEmail(iEmail.Sender, iEmail.Subject, iEmail.Content, iEmail.Recipients)
	if err != nil {
		t.Error("hermes:NewEmail()", err)
	}
	if err = emailProvider.Send(email); err != nil {
		t.Error("SendEmail()", err)
	}

	return
}

func TestSendEmailFail(t *testing.T) {
	//TODO
	return
}

func TestChangeProviderOK(t *testing.T) {
	p1Name := emailProvider.Selected()
	emailProvider.NextProvider()
	p2Name := emailProvider.Selected()
	log.Println("Provder1Name:", p1Name)
	log.Println("Provder1Name:", p2Name)
	if p1Name == p2Name {
		t.Error("Wrong change from provider")
	}

	return
}

func TestChangeProviderFail(t *testing.T) {
	//TODO: Implement
	return
}

func TestSortProvidersOK(t *testing.T) {
	log.Println("Order:", emailProvider.Order())
	newOrder := []string{"sendgrid", "mailgun", "mailchimp"}
	emailProvider.Sort(newOrder...)
	log.Println("New Order:", emailProvider.Order())

	return
}
