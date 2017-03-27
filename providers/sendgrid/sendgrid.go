package sendgrid

import (
	"fmt"
	"errors"
	"github.com/mauricio-cdr/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	cfgfile     = `provider.conf`
	sApiVersion = "/v3/mail/send"
	sBasePath   = "https://api.sendgrid.com"
)

type Sendgrid struct {
	ID        int64  `json:"-" db:"id"`
	Name      string `json:"-" db:"provider_name"`
	APIKey    string `json:"-" db:"-"`
	ConunterM int64  `json:"-" db:"counter_m"`
}

func NewProvider() *Sendgrid {
	s := &Sendgrid{}
	return s
}

func (p *Sendgrid) GetName() (name string) {
	return `sendgrid`
}

func (p *Sendgrid) Init() (err error) {
	c, err := config.NewConfig(cfgfile)
	p.Name = p.GetName()
	if p.APIKey, err = c.Property(p.Name, "apikey"); err != nil {
		return errors.New("ERR_INVALID_APIKEY")
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
		fmt.Printf("----------------------%+v",err)
		return errors.New("ERR_INVALID_MESSAGE")
	} else {
		if response.StatusCode == 429 {
			return errors.New("ERR_LIMIT_REACHED")
		} else if response.StatusCode == 401 {
			return errors.New("ERR_INVALID_APIKEY")
		}
		return
	}
}

func (p *Sendgrid) NewEmail(sender interface{},s string, c interface{}) (ms interface{}, err error) {
	var m = Email{}
	m = NewEmail()
	m.AddSender(sender)
	m.AddSubject(s)
	m.AddContent(c)
	ms = &m
	return
}

func (p *Sendgrid) ToString() string {
	return "Name:" + p.Name + "-APIKey:" + p.APIKey
}
