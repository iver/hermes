package mailchimp

import (
	"os"
	"path"
	"fmt"
	"github.com/mattbaird/gochimp"
	"bitbucket.org/ivan-iver/config"
	"errors"
)

var (
	cfgfile   = `provider.conf`
)

// Email interface
type Email interface {
   AddSenderEmail(e string) error
   AddSubject(s string) error
   AddSenderName(name string) error
   AddRecipients(e ...string) error
   AddAttachment(p string) error
   AddTemplate(t string) error
   AddContent(c string) error
}


type Mailchimp struct {
	 ID           int64
     Name 		  string 
	 APIKey       string
	 CounterM     int
	 MandrillAPI  *gochimp.MandrillAPI
}

func NewProvider() *Mailchimp {
    s := &Mailchimp{}
    return s
}

func (p *Mailchimp) GetName() (name string) {
	return `mailchimp`
}

func (p *Mailchimp) Init() (err error) {
    c,err := Config()
    p.APIKey,err = c.String("mailchimp", "apikey") 
    if err!=nil{
       return errors.New("ERR_INVALID_APIKEY")
	}
	p.Name = p.GetName()
    if p.MandrillAPI, err = gochimp.NewMandrill(p.APIKey); err!=nil{
       return errors.New("ERR_INVALID_APIKEY")
	}
	return
}

	
//  sendemail function with mailchimp provider
func (p *Mailchimp) SendEmail(email Email) (err error) {
	var emailg = email.(*MEmail)
    response, err := p.MandrillAPI.MessageSend(emailg.GochimpM, false);
	if err != nil {
		return errors.New("ERR_INVALID_MESSAGE")
	}
 
    fmt.Printf("%+v",response)
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

func (p *Mailchimp) NewEmail(se string , sn string , s string ,t string) (m Email,err error) {
	 var mm =MEmail{}
	 mm.AddSenderEmail(se)
	 mm.AddSenderName(sn)
	 mm.AddSubject(s)
	 mm.AddTemplate(t)
	 m= &mm
     return 
}