package mailchimp

import (
	"fmt"
	"os"
	"github.com/mattbaird/gochimp"
	"errors"
)

type Mailchimp struct {
	 APIKey   string
	 CounterM int
	 MandrillAPI *gochimp.MandrillAPI 
}

func (p *Mailchimp) Init() (err error) {
    p.APIKey = os.Getenv("MANDRILL_KEY")
    p.MandrillAPI, err = gochimp.NewMandrill(p.APIKey)
	if p.APIKey == ""{
       return errors.New("ERR_INVALID_APIKEY")
	}
	return
}


//  sendemail function with mailchimp provider
func (p *Mailchimp) SendEmail(email MailchimpEmail) (err error) {
    response, err := p.MandrillAPI.MessageSend(email.GochimpM, false);
	if err != nil {
		return errors.New("ERR_INVALID_MESSAGE")
	}
 
    fmt.Printf("%+v",response)
	return
}

func (p *Mailchimp) NewEmail(se string , sn string , s string ,t string) (m MailchimpEmail,err error) {
	 m.AddSenderEmail(se)
	 m.AddSenderName(sn)
	 m.AddSubject(s)
	 m.AddTemplate(t)
     return 
}