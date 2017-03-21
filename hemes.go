package hermes

import (
	"container/ring"
	"github.com/ivan-iver/hermes/lib"
	"github.com/ivan-iver/hermes/providers"
	"errors"
)

var (
	DefaultOrder = []string{"mailchimp", "mailgun", "sendgrid"}
)

type EmailProvider struct {
	Providers *ring.Ring
}

func New() (e EmailProvider, err error) {
	var invalidProviders = 0
	e = EmailProvider{}
	e.Providers = ring.New(len(DefaultOrder))
	for i := 0; i < e.Providers.Len(); i++ {
		if p, err := providers.NewProvider([]string{DefaultOrder[i]}); err == nil {
			provider:=  p.(lib.Provider)
			provider.Init()
			e.Providers.Value = provider
			
			e.Providers = e.Providers.Next()
		}else{
           invalidProviders++
		}
	}
	if invalidProviders == len(DefaultOrder){
		err = errors.New("ERR_INVALID_PROVIDERS")
	}
	return	
}

func (e *EmailProvider) NextProvider() {
	e.Providers = e.Providers.Next()
}

func  (e *EmailProvider) Order() string{
	order:=""
	e.Providers.Do(func (p interface{}){
		provider := p.(lib.Provider)
		order += provider.GetName()+" " 
	})
	return order
}

func (e *EmailProvider) Sort(order ...string) (err error) {
     var invalidProviders = 0
	for i := 0; i < e.Providers.Len(); i++ {
		if p, err := providers.NewProvider([]string{order[i]}); err == nil {
			provider:=  p.(lib.Provider)
			provider.Init()
			e.Providers.Value = provider
			
			e.Providers = e.Providers.Next()
		}else{
           invalidProviders++
		}
	}
	if invalidProviders == len(DefaultOrder){
		err = errors.New("ERR_INVALID_PROVIDERS")
	}
	return	
}

func (e *EmailProvider) Send(m interface{}) (err error) {
	var emailSended = false
	head := e.Providers.Value.(lib.Provider)	
	for !emailSended {
	   provider := e.Providers.Value.(lib.Provider)	
       if err = provider.SendEmail(m); err != nil {
		 if err.Error() == "ERR_LIMIT_REACHED"{
            e.Providers = e.Providers.Next()
			current := e.Providers.Value.(lib.Provider)
			if providers.Equals(head,current){
				return errors.New("ERR_ALL_LIMITS_REACHED")
			}
		 }
	   }else{
		  emailSended=true 
	   }
	}
	return
}

func (e *EmailProvider) SelectedProvider() (pn string) {
	provider := e.Providers.Value.(lib.Provider)
	pn = provider.GetName()
	return
}

func (e *EmailProvider) NewEmail(d string, sn string, s string, t string, r ...string) (email interface{},err error) {
	provider := e.Providers.Value.(lib.Provider)
    emailI,err := provider.NewEmail(d,sn,s,t)
	emailp := emailI.(lib.Email)
	emailp.AddRecipients(r...)
	email= emailp
	return
}
