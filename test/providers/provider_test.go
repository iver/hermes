package dd

import (
    "log"
	"testing"
    "github.com/ivan-iver/hermes/providers"
	"github.com/ivan-iver/hermes/models"
    )  

func TestCreateProvider(t *testing.T) {
	var options = []string{"mailchimp"}
    var err error
	var providerI interface{}
	if providerI,err = providers.NewProvider(options); err != nil {
		t.Error("providers:NewProvider()", err)
	}
	provider:=providerI.(models.Provider)
    log.Printf("Provider name:%+v",provider)
	return

}
