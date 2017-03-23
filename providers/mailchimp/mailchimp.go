package mailchimp

import (
	"errors"
	"github.com/mattbaird/gochimp"
	"github.com/mauricio-cdr/config"
)

var (
	cfgfile = `provider.conf`
)

type Mailchimp struct {
	ID          int64
	Name        string
	APIKey      string
	CounterM    int
	MandrillAPI *gochimp.MandrillAPI
}

func NewProvider() *Mailchimp {
	s := &Mailchimp{}
	return s
}

func (p *Mailchimp) GetName() (name string) {
	return `mailchimp`
}

func (p *Mailchimp) Init() (err error) {
	c, err := config.NewConfig(cfgfile)
	p.Name = p.GetName()
	p.APIKey, err = c.Property(p.Name, "apikey")
	if err != nil {
		return errors.New("ERR_INVALID_APIKEY")
	}
	if p.MandrillAPI, err = gochimp.NewMandrill(p.APIKey); err != nil {
		return errors.New("ERR_INVALID_APIKEY")
	}
	return
}

//  sendemail function with mailchimp provider
func (p *Mailchimp) SendEmail(emailI interface{}) (err error) {
	email := *emailI.(*Email)
	_, err = p.MandrillAPI.MessageSend(email.GochimpM, false)
	if err != nil {
		return errors.New("ERR_INVALID_MESSAGE")
	}
	return
}

func (p *Mailchimp) NewEmail(se interface{}, s string,t interface{}) (m interface{}, err error) {
	var mm = Email{}
	mm.AddSender(se)
	mm.AddSubject(s)
	mm.AddTemplate(t)
	m = &mm
	return
}

func (p *Mailchimp) ToString() string {
	return "Name:" + p.Name + "-APIKey:" + p.APIKey
}
