package mailgun

import (
	"os"
    "log"
	"errors"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type Mailgun struct {
     Domain       string
	 APIKey       string
	 PublicAPIKey string
	 CounterM     int
}

func (p *Mailgun) Init() (err error) {
	 p.Domain = os.Getenv("MG_DOMAIN")
	 p.APIKey = os.Getenv("MG_API_KEY")
	 p.PublicAPIKey = os.Getenv("MG_PUBLIC_API_KEY")
	 if p.Domain == "" {return errors.New("ERR_INVALID_DOMAIN")}
	 if p.APIKey == "" {return errors.New("ERR_INVALID_APIKEY")}
	 if p.PublicAPIKey == "" {return errors.New("ERR_INVALID_PUBAPIKEY")}
	 return
}

// send email method with mailgun provider

func (p *Mailgun) SendEmail(email MailgunEmail) (err error) {

	mg := mailgun.NewMailgun(p.Domain, p.APIKey, p.PublicAPIKey)
	
	message := mailgun.NewMessage(
    "mailgun@hermes.mx",
    "Un saludo",
    "<div><h1>Hola desde mailgun<h1><h4>Template desde hermes</h4></div>",
    "mau.cdr.19@gmail.com")
	
	resp, id, err := mg.Send(message)
	if err != nil {
	  return errors.New("ERR_INVALID_MESSAGE")
	}
	log.Printf("--------%+v,%+v",resp,id)

	return
}


func (p *Mailgun) NewEmail(se string , sn string , s string ,t string) (m MailgunEmail,err error) {
	 m.AddSubject(s)
	 m.AddSenderEmail(se)
	 m.AddContent(t)
	 m.AddSenderName(sn)
	 m.SetValues()
     return 

}



