package providers_test

import (
    "log"
	"testing"
    "github.com/ivan-iver/hermes/providers"
	"github.com/ivan-iver/hermes/lib"
    )  

func TestCreateProviders(t *testing.T) {
	var err error
	var options1 = []string{"mailchimp"}
	var options2 = []string{"sendgrid"}
	var options3 = []string{"mailgun"}
    var providerI interface{}
	
	if providerI,err = providers.NewProvider(options1); err != nil {
		t.Error("CreateProviders:NewProvider()", err)
	}
	mailchimp:=providerI.(lib.Provider)

    if name:=mailchimp.GetName(); name!= "mailchimp" {
		log.Printf("ProvderName:%v",name)
		t.Error("GetName:Diferent to mailchimp", err)
	}
	if providerI,err = providers.NewProvider(options2); err != nil {
		t.Error("CreateProviders:NewProvider()", err)
	}
	sendgrid:=providerI.(lib.Provider)

    if name :=sendgrid.GetName(); name != "sendgrid" {
		log.Printf("ProvderName:%v",name)
		t.Error("GetName:Diferent to sendgrid", err)
	}
	if providerI,err = providers.NewProvider(options3); err != nil {
		t.Error("CreateProviders:NewProvider()", err)
	}
	mailgun:=providerI.(lib.Provider)

    if name:=mailgun.GetName(); name != "mailgun" {
		log.Printf("ProvderName:%v",name)
		t.Error("GetName:Diferent to mailgun", err)
	}
	return

}
