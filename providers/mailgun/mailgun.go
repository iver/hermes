package mailgun

import (
	"os"
	"fmt"
	"path"
	"errors"
	"gopkg.in/mailgun/mailgun-go.v1"
	"bitbucket.org/ivan-iver/config"
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

func NewMailgun() *Mailgun {
    s := &Mailgun{}
    return s
}

func (p *Mailgun) GetName() (name string) {
	return `mailgun`
}


func (p *Mailgun) Init() (err error) {
	 c,err := Config()
	 p.Name = p.GetName()
	 if p.PublicAPIKey,err = c.String("mailgun", "publicapikey");err!=nil{
		return errors.New("ERR_INVALID_PUBAPIKEY")
	 }
	 if p.APIKey,err = c.String("mailgun", "apikey");err!=nil{
		return errors.New("ERR_INVALID_APIKEY")
	 }
	 //if p.Domain,err = c.String("mailgun", "domain");err!=nil{
		p.Domain="sandboxeceaa78856bd40e1bee2496e06f723b0.mailgun.org"
	 //}
	 
	 return
}

// send email method with mailgun provider

func (p *Mailgun) SendEmail(email Email) (err error) {

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

func Config() (cfg *config.Config,err error){ 
	var pwd string
  	if pwd, err = os.Getwd(); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}
    pwd = path.Join(pwd, cfgfile)
	if cfg, err = config.ReadDefault(pwd); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}
	return
}

func (p *Mailgun) NewEmail(se string , sn string , s string ,t string) (m Email,err error) {
	 m.AddSubject(s)
	 m.AddSenderEmail(se)
	 m.AddContent(t)
	 m.AddSenderName(sn)
	 m.SetValues()
     return 

}



