package hermes

import (
	"container/ring"
	"github.com/ivan-iver/hermes/lib"
	"github.com/ivan-iver/hermes/providers"
)

var (
	DefaultOrder = []string{"mailchimp", "mailgun", "sendgrid"}
)

type EmailProvider struct {
	Providers *ring.Ring
}

func New() (e EmailProvider, err error) {
	e = EmailProvider{}

	e.Providers = ring.New(len(DefaultOrder))
	for i := 0; i < e.Providers.Len(); i++ {
		if p, err := providers.NewProvider([]string{DefaultOrder[i]}); err == nil {
			e.Providers.Value = p.(lib.Provider)
			e.Providers = e.Providers.Next()
		}

	}
	return
}

func (e *EmailProvider) NextProvider() {
	e.Providers = e.Providers.Next()
}

func (e *EmailProvider) Sort(order ...string) (err error) {
	 newOrder := ring.New(len(order))

	 for i := 0; i < newOrder.Len(); i++ {
		if p, err := providers.NewProvider([]string{order[i]}); err == nil {
			e.Providers.Value = p.(lib.Provider)
			e.Providers = e.Providers.Next()
		}

	}
	return
}

func (e *EmailProvider) Send(m *lib.Email) (err error) {
	provider := e.Providers.Value.(lib.Provider)
	if err = provider.SendEmail(m); err != nil {
		 if err.Error()== "ERR_LIMIT_REACHED"{
        e.Providers = e.Providers.Next()
				//TODO add validation with all providers
		 }
	}
	return
}

func (e *EmailProvider) SelectedProvider() (pn string) {
	provider := e.Providers.Value.(lib.Provider)
	pn = provider.GetName()
	return
}

func (e *EmailProvider) NewEmail(d string, sn string, s string, t string, r ...string) (email lib.Email,err error) {
	provider := e.Providers.Value.(lib.Provider)
  emailI,err := provider.NewEmail(d,sn,s,t)
	email = emailI.(lib.Email)
	email.AddRecipients(r...)
	return
}
