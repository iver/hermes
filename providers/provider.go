package providers

import (
	"github.com/ivan-iver/hermes/lib"
	"github.com/ivan-iver/hermes/models"
	"github.com/ivan-iver/hermes/providers/mailchimp"
	"github.com/ivan-iver/hermes/providers/mailgun"
	"github.com/ivan-iver/hermes/providers/sendgrid"
)

func NewProvider(cfile string, options []string) (p interface{}, err error) {
	var first = options[0]
	switch first {
	case "mailgun":
		p = mailgun.NewProvider(cfile)
	case "mailchimp":
		p = mailchimp.NewProvider(cfile)
	case "sendgrid":
		p = sendgrid.NewProvider(cfile)
	default:
		err = models.ErrUnknownProvider
	}
	return
}

func Equals(providerI1 interface{}, providerI2 interface{}) bool {
	var ok bool
	var provider1 lib.Provider
	var provider2 lib.Provider
	if provider1, ok = providerI1.(lib.Provider); !ok {
		return false
	}

	if provider2, ok = providerI2.(lib.Provider); !ok {
		return false
	}

	return (provider1.ToString() == provider2.ToString())

}
