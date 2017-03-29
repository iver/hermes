package mailchimp

import (
	"github.com/mattbaird/gochimp"
	"github.com/mauricio-cdr/config"
	"github.com/ivan-iver/hermes/models"
)

var (
	Cfgfile = `provider.conf`
)

type Mailchimp struct {
	ID          int64                 `json:"id,omitempty"`
	Name        string                `json:"name,omitempty"`
	APIKey      string                `json:"api_key,omitempty"`
	MandrillAPI *gochimp.MandrillAPI  `json:"mandril_api,omitempty"`
	CounterM    int64                 `json:"counter_m,omitempty"`
}

func NewProvider() *Mailchimp {
	s := &Mailchimp{}
	return s
}

func (p *Mailchimp) GetName() (name string) {
	return `mailchimp`
}

func (p *Mailchimp) Init() (err error) {
	c, err := config.NewConfig(Cfgfile)
	p.Name = p.GetName()
	p.APIKey, err = c.Property(p.Name, "apikey")
	if err != nil {
		return models.ErrInvalidAPIKey
	}
	if p.MandrillAPI, err = gochimp.NewMandrill(p.APIKey); err != nil {
		return models.ErrInvalidAPIKey
	}
	return
}

//  SendEmail function with mailchimp provider
func (p *Mailchimp) SendEmail(emailI interface{}) (err error) {
	email := *emailI.(*Email)
	_, err = p.MandrillAPI.MessageSend(email.GochimpM, false)
	if err != nil {
		return models.ErrInvalidMessage
	}
	return
}

func (p *Mailchimp) NewEmail(se interface{}, s string,c interface{}) (m interface{}, err error) {
	var mm = Email{}
	if err=mm.AddSender(se);err!=nil{
		return m,err
	}
	if err=mm.AddSubject(s);err!=nil{
		return 
	}
	if err=mm.AddContent(c);err!=nil{
		return 
	}

	m = &mm
	return
}

func (p *Mailchimp) RefactorEmail(mail map[string]interface{})(ms interface{}, err error){
	var m = Email{}
	m.AddSender(mail["sender"])
	m.AddSubject(mail["subject"].(string))
	m.AddContent(mail["content"])
	m.AddRecipients(mail["recipients"])
	ms = &m
	return
}
func (p *Mailchimp) ToString() string {
	return "Name:" + p.Name + "-APIKey:" + p.APIKey
}
