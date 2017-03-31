package mailgun

import (
	"github.com/notifik/config"
	"gopkg.in/mailgun/mailgun-go.v1"
	"github.com/ivan-iver/hermes/models"
)

var (
	Cfgfile = `provider.conf`
)

type Mailgun struct {
	ID           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Domain       string `json:"domain,omitempty"`
	APIKey       string `json:"api_key,omitempty"`
	PublicAPIKey string `json:"public_api_key,omitempty"`
	CounterM     int64  `json:"counter_m,omitempty"`
}

func NewProvider() *Mailgun {
	s := &Mailgun{}
	return s
}

func (p *Mailgun) GetName() (name string) {
	return `mailgun`
}

func (p *Mailgun) Init() (err error) {
	c, err := config.NewConfig(Cfgfile)
	p.Name = p.GetName()
	if p.PublicAPIKey, err = c.Property(p.Name, "publicapikey"); err != nil {
		return models.ErrInvalidPublicAPIKey
	}
	if p.APIKey, err = c.Property(p.Name, "apikey"); err != nil {
		return models.ErrInvalidAPIKey
	}
	if p.Domain, err = c.Property(p.Name, "domain"); err != nil {
		return models.ErrInvalidDomain
	}

	return
}

// send email method with mailgun provider

func (p *Mailgun) SendEmail(emailI interface{}) (err error) {
	mg := mailgun.NewMailgun(p.Domain, p.APIKey, p.PublicAPIKey)

	message := mailgun.NewMessage(
		"mailgun@hermes.mx",
		"Un saludo",
		"<div><h1>Hola desde mailgun<h1><h4>Template desde hermes</h4></div>",
		"mau.cdr.19@gmail.com")

	_, _, err = mg.Send(message)
	if err != nil {
		return models.ErrInvalidMessage
	}
	return
}

func (p *Mailgun) NewEmail(se interface{}, s string, c interface{}) (ms interface{}, err error) {
	var m = Email{}
	if err=m.AddSender(se);err!=nil{
		return m,err
	}
	if err=m.AddSubject(s);err!=nil{
		return 
	}
	if err=m.AddContent(c);err!=nil{
		return 
	}
	m.SetValues()
	ms = &m
	return

}

func (p *Mailgun) RefactorEmail(mail map[string]interface{})(ms interface{}, err error){
	var m = Email{}
	m.AddSender(mail["sender"])
	m.AddSubject(mail["subject"].(string))
	m.AddContent(mail["content"])
	m.SetValues()
	m.AddRecipients(mail["recipients"])
	ms = &m
	return
}

func (p *Mailgun) ToString() string {
	return "Name:" + p.Name + "-APIKey:" + p.APIKey
}
