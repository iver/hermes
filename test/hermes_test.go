package hermes_test

import (
	 "log"
	 "testing"
	 "github.com/ivan-iver/hermes"
)

var (
	emailProvider hermes.EmailProvider
)


func TestCreateHermesOK(t *testing.T) {
    var err error
	if emailProvider,err=hermes.New(); err!=nil{
       t.Error("hermes:New()", err)
	}
	if name:=emailProvider.SelectedProvider(); name !="mailchimp"{
	   log.Println("ProviderSelected:",name)	
	   t.Error("Wrong provider")
	}
	return
}

func TestSendEmailOK(t *testing.T) {
    var err error
	iEmail:= CorrectEmail()
    email,err:=emailProvider.NewEmail(iEmail.SenderEmail,iEmail.SenderName,iEmail.Subject,iEmail.Content,iEmail.Recipients...); 
	if err!=nil{
       t.Error("hermes:NewEmail()", err)
	} 
	if err = emailProvider.Send(email); err!=nil{	
	  t.Error("SendEmail()",err)
	}
	
	return
}

func TestChangeProviderOK(t *testing.T) {
	p1Name:= emailProvider.SelectedProvider()
	emailProvider.NextProvider()
	p2Name:= emailProvider.SelectedProvider()
	log.Println("Provder1Name:",p1Name)
	log.Println("Provder1Name:",p2Name)
	if p1Name == p2Name{	
	  t.Error("Wrong change from provider")
	}
	
	return
}

func TestSortProvidersOK(t *testing.T) {
	log.Println("Order:",emailProvider.Order())
	newOrder:= []string{"sendgrid","mailgun","mailchimp"}
	emailProvider.Sort(newOrder...)
	log.Println("New Order:",emailProvider.Order())
	
	return
}
