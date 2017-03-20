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
func (p *Mailchimp) SendEmail(emailI interface{}) (err error) {
	email:=emailI.(*Email)
    response, err := p.MandrillAPI.MessageSend(email.GochimpM, false);
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

func (p *Mailchimp) NewEmail(se string , sn string , s string ,t string) (m interface{},err error) {
	 var mm =Email{}
	 mm.AddSenderEmail(se)
	 mm.AddSenderName(sn)
	 mm.AddSubject(s)
	 mm.AddTemplate(t)
	 m= &mm
     return 
}

func (p *Mailchimp) ToString() string{
	 return "Name:"+p.Name+"-APIKey:"+p.APIKey
}