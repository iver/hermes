package hermes

import (
	"container/ring"

	"github.com/ivan-iver/hermes/lib"
	"github.com/ivan-iver/hermes/models"
	"github.com/ivan-iver/hermes/providers"
)

var (
	DefaultOrder = []string{"mailchimp", "mailgun", "sendgrid"}
)

type EmailProvider struct {
	ConfigFile string
	Providers *ring.Ring
}

func New(cfile string) (e *EmailProvider, err error) {
	var invalidProviders = 0
	e = &EmailProvider{ConfigFile:cfile}
	e.Providers = ring.New(len(DefaultOrder))
	for i := 0; i < e.Providers.Len(); i++ {
		if p, err := providers.NewProvider(cfile, []string{DefaultOrder[i]}); err == nil {
			provider := p.(lib.Provider)
			provider.Init()
			e.Providers.Value = provider

			e.Providers = e.Providers.Next()
		} else {
			invalidProviders++
		}
	}
	if invalidProviders == len(DefaultOrder) {
		err = models.ErrInvalidProviders
	}
	return
}

func (e *EmailProvider) NextProvider() {
	e.Providers = e.Providers.Next()
}

func (e *EmailProvider) Order() string {
	order := ""
	e.Providers.Do(func(p interface{}) {
		provider := p.(lib.Provider)
		order += provider.GetName() + " "
	})
	return order
}

func (e *EmailProvider) Sort(order ...string) (err error) {
	var invalidProviders = 0
	for i := 0; i < e.Providers.Len(); i++ {
		if p, err := providers.NewProvider(e.ConfigFile,[]string{order[i]}); err == nil {
			provider := p.(lib.Provider)
			provider.Init()
			e.Providers.Value = provider

			e.Providers = e.Providers.Next()
		} else {
			invalidProviders++
		}
	}
	if invalidProviders == len(DefaultOrder) {
		err = models.ErrInvalidProviders
	}
	return
}

//Send ...
func (e *EmailProvider) Send(m interface{}) (err error) {
	var emailSended = false
	var emailToSend = m.(lib.Email)
	head := e.Providers.Value.(lib.Provider)
	for !emailSended {
		provider := e.Providers.Value.(lib.Provider)
		if err = provider.SendEmail(emailToSend); err != nil {
			if err == models.ErrLimitMessagesReached {
				email, _ := provider.RefactorEmail(emailToSend.GetInfo())
				emailToSend = email.(lib.Email)
				e.NextProvider()
				current := e.Providers.Value.(lib.Provider)
				if providers.Equals(head, current) {
					return models.ErrAllLimitsReached
				}
			} else {
				return err
			}
		} else {
			emailSended = true
		}
	}
	return
}

func (e *EmailProvider) Selected() (pn string) {
	provider := e.Providers.Value.(lib.Provider)
	pn = provider.GetName()
	return
}

func (e *EmailProvider) NewEmail(sn interface{}, s string, t interface{}, r interface{}) (email interface{}, err error) {
	provider := e.Providers.Value.(lib.Provider)
	emailI, err := provider.NewEmail(sn, s, t)
	emailp := emailI.(lib.Email)
	emailp.AddRecipients(r)
	email = emailp
	return
}
