package sendgrid

import (
	"os"
	"fmt"
	"path"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go"
	"bitbucket.org/ivan-iver/config"
	"errors"
)

var (
	cfgfile    = `provider.conf`
	sApiVersion = "/v3/mail/send"
	sBasePath = "https://api.sendgrid.com"
)

type Sendgrid struct {
	 ID           int64        `json:"-" db:"id"`
     Name         string       `json:"-" db:"provider_name"`
	 APIKey       string       `json:"-" db:"-"`
	 ConunterM    int64        `json:"-" db:"counter_m"`
}

func NewProvider() *Sendgrid {
    s := &Sendgrid{}
    return s
}

func (p *Sendgrid) GetName() (name string) {
	return `sendgrid`
}

func (p *Sendgrid) Init() (err error) {
	 c,err := Config()
	p.Name = p.GetName()
   	if p.APIKey,err = c.String("sendgrid", "apikey"); err !=nil{
       return errors.New("ERR_INVALID_APIKEY")
	}
	return
}

//  sendemail function with sendgrid provider
func (p *Sendgrid) SendEmail(emailI interface{}) (err error) {
	email:=emailI.(*Email)
	request := sendgrid.GetRequest(p.APIKey,sApiVersion,sBasePath)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(email.SendgridM)
	response, err := sendgrid.API(request)
	if err != nil {
		return errors.New("ERR_INVALID_MESSAGE")
    } else {
        if response.StatusCode == 429 {
             return errors.New("ERR_LIMIT_REACHED")
		}else if response.StatusCode == 401 {
			 return errors.New("ERR_INVALID_APIKEY")
		}
		return 
    }
}

func (p *Sendgrid) NewEmail(se string , sn string , s string ,t string) (ms interface{},err error) {
	 var m =Email{}
	 m=NewEmail()
	 m.AddSubject(s)
	 m.AddSenderEmail(se)
	 m.AddContent(t)
	 m.AddSenderName(sn)
	 ms =&m
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
