package mailchimp

import (
	"fmt"
	"os"
	"github.com/mattbaird/gochimp"
)

type Mailchimp struct {
	 APIKey string
}

func (p *Mailchimp) Init() (err error) {
     p.APIKey = os.Getenv("MANDRILL_KEY")
	  if err != nil {
        fmt.Println("Error instantiating client")
    }
	return
}


//  sendemail function with mailchimp provider
func (p *Mailchimp) SendEmail(email MailchimpEmail) (err error) {

    mandrillAPI, err := gochimp.NewMandrill(p.APIKey)
	if err != nil {
		return err
	}
	if _, err = mandrillAPI.MessageSend(email.GochimpM, false); err != nil {
		fmt.Println("Error sending message")
		return err
	}

	return
}

func (p *Mailchimp) NewEmail(se string , sn string , s string ,t string) (m MailchimpEmail,err error) {
	 m.AddSenderEmail(se)
	 m.AddSenderName(sn)
	 m.AddSubject(s)
	 m.AddTemplate(t)
     return 
}