package sendgrid

import (
	"os"
    "log"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go"
	"errors"
)

type Sendgrid struct {
	 APIKey string
}

func (p *Sendgrid) Init() (err error) {
   	p.APIKey = os.Getenv("SENDGRID_API_KEY")
	return
}


//  sendemail function with sendgrid provider
func (p *Sendgrid) SendEmail(email SendgridEmail) (err error) {
	request := sendgrid.GetRequest(p.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(email.SendgridM)
	response, err := sendgrid.API(request)
	log.Println(response,err)
	if err != nil {
        log.Println(err)
		return err
    } else {
		
        if response.StatusCode == 429 {
             return errors.New(`error sending message`)
		}
		  return response
    }
}

func (p *Sendgrid) NewEmail(se string , sn string , s string ,t string) (m SendgridEmail,err error) {
	 m=NewEmail()
	 m.AddSubject(s)
	 m.AddSenderEmail(se)
	 m.AddContent(t)
	 m.AddSenderName(sn)
     return 
}