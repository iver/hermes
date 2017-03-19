package hermes

import (
	"github.com/ivan-iver/hermes/providers"
	"github.com/ivan-iver/hermes/lib"
)

var (
	 DefaultOrder = []string{ "mailchimp", "mailgun", "sendgrid"}
)
type EmailProvider struct {
	 Providers map[string]lib.Provider
	 OrderProviders []string
	 ActualProvider string
}

func New() (e EmailProvider,err error){
    e = EmailProvider{}
	  e.Providers = map[string]lib.Provider{ "mailchimp": nil,
                   					   "mailgun"  : nil,
                    				   "sendgrid" : nil,
    }
		e.OrderProviders = DefaultOrder
    e.ActualProvider=e.OrderProviders[0]
	  for key, _ := range e.Providers {
		   if p,err:= providers.NewProvider([]string{key}); err==nil {
              e.Providers[key] = p.(lib.Provider)
		   }
    }
	return 
}


func (e *EmailProvider)Sort(order ...string) (err error){
  e.OrderProviders=order
  return 
}

func (e *EmailProvider) Send(m *lib.Email)(err error){
  return
}

func (e *EmailProvider) SelectedProvider()(pn string){
	pn= e.Providers[e.ActualProvider].GetName()
  return
}

func (p *EmailProvider) NewEmail(e string , sn string , s string ,t string,r ...string) (err error){
  return
}