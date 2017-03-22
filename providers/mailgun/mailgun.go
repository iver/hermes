package mailgun

import (
	"errors"
	"gopkg.in/mailgun/mailgun-go.v1"
	"github.com/mauricio-cdr/config"
)

var (
	cfgfile   = `provider.conf`
)

type Mailgun struct {
	 ID           int64        `json:"-" db:"id"`
     Name 		  string       `json:"-" db:"provider_name"`
     Domain       string       `json:"-" db:"-"`
	 APIKey       string       `json:"-" db:"-"`
	 PublicAPIKey string       `json:"-" db:"-"`
	 CounterM     int64        `json:"-" db:"counter"`
}

func NewProvider() *Mailgun {
    s := &Mailgun{}
    return s
}

func (p *Mailgun) GetName() (name string) {
	return `mailgun`
}


func (p *Mailgun) Init() (err error) {
	 c,err := config.NewConfig(cfgfile)
	 p.Name = p.GetName()
	 if p.PublicAPIKey,err = c.Property(p.Name, "publicapikey");err!=nil{
		return errors.New("ERR_INVALID_PUBAPIKEY")
	 }
	 if p.APIKey,err = c.Property(p.Name, "apikey");err!=nil{
		return errors.New("ERR_INVALID_APIKEY")
	 }
	 if p.Domain,err = c.Property(p.Name, "domain");err!=nil{
		//p.Domain="sandboxeceaa78856bd40e1bee2496e06f723b0.mailgun.org"
		return errors.New("ERR_INVALID_APIKEY")
	 }
	 
	 return
}

// send email method with mailgun provider

func (p *Mailgun) SendEmail(emailI interface{}) (err error) {
    //email:=emailI.(Email)
	mg := mailgun.NewMailgun(p.Domain, p.APIKey, p.PublicAPIKey)
	
	message := mailgun.NewMessage(
    "mailgun@hermes.mx",
    "Un saludo",
    "<div><h1>Hola desde mailgun<h1><h4>Template desde hermes</h4></div>",
    "mau.cdr.19@gmail.com")
	
	_, _, err = mg.Send(message)
	if err != nil {
	  return errors.New("ERR_INVALID_MESSAGE")
	}
	return
}

func (p *Mailgun) NewEmail(se string , sn string , s string ,t string) (ms interface{} ,err error) {
     var m =Email{}
	 m.AddSubject(s)
	 m.AddSenderEmail(se)
	 m.AddContent(t)
	 m.AddSenderName(sn)
	 m.SetValues()
	 ms= &m
     return 

}


func (p *Mailgun) ToString() string{
	 return "Name:"+p.Name+"-APIKey:"+p.APIKey
}

