package sendgrid

import (
	"github.com/mauricio-cdr/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/ivan-iver/hermes/models"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	Cfgfile     = `provider.conf`
	sApiVersion = "/v3/mail/send"
	sBasePath   = "https://api.sendgrid.com"
)

type Sendgrid struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	APIKey    string `json:"api_key,omitempty"`
	ConunterM int64  `json:"counter_m,omitempty"`
}

func NewProvider() *Sendgrid {
	s := &Sendgrid{}
	return s
}

func (p *Sendgrid) GetName() (name string) {
	return `sendgrid`
}

func (p *Sendgrid) Init() (err error) {
	c, err := config.NewConfig(Cfgfile)
	p.Name = p.GetName()
	if p.APIKey, err = c.Property(p.Name, "apikey"); err != nil {
		return models.ErrInvalidAPIKey
	}
	return
}

//  sendemail function with sendgrid provider
func (p *Sendgrid) SendEmail(emailI interface{}) (err error) {
	email := emailI.(*Email)
	request := sendgrid.GetRequest(p.APIKey, sApiVersion, sBasePath)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(email.SendgridM)
	response, err := sendgrid.API(request)
	if err != nil {
		return models.ErrInvalidMessage
	} else {
		if response.StatusCode == 429 {
			return models.ErrLimitMessagesReached
		} else if response.StatusCode == 401 {
			return models.ErrInvalidAPIKey
		}
		return
	}
}

func (p *Sendgrid) NewEmail(se interface{},s string, c interface{}) (ms interface{}, err error) {
	var m = Email{}
	m = NewEmail()
	if err=m.AddSender(se);err!=nil{
		return m,err
	}
	if err=m.AddSubject(s);err!=nil{
		return 
	}
	if err=m.AddContent(c);err!=nil{
		return 
	}
	ms = &m
	return
}

func (p *Sendgrid) RefactorEmail(mail map[string]interface{})(ms interface{}, err error){
	var m = Email{}
	m = NewEmail()
	m.AddSender(mail["sender"])
	m.AddSubject(mail["subject"].(string))
	m.AddContent(mail["content"])
	m.AddRecipients(mail["recipients"])
	ms = &m
	return
}

func (p *Sendgrid) ToString() string {
	return "Name:" + p.Name + "-APIKey:" + p.APIKey
}
