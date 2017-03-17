package mailchimp
 import (
	 "time"
	 "os"
	 "fmt"
	 "path"
	 "github.com/mattbaird/gochimp"
	 "bitbucket.org/ivan-iver/config"
 ) 

type MEmail struct{
	ID          int64          `json:"-" db:"id,omitempty"`
	InsertedAt  time.Time      `json:"-" db:"inserted_at,omitempty"`
	SendedAt    time.Time      `json:"-" db:"sended_at,omitempty"`  
	GochimpM    gochimp.Message`json:"-" db:"gochimp_message,omitempty"`
}

func (m *MEmail) AddSenderEmail(e string) (err error){
   m.GochimpM.FromEmail=e
   return
}

func (m *MEmail) AddSenderName(n string) (err error){
   m.GochimpM.FromName= n
   return
}

func (m *MEmail) AddSubject(s string) (err error){
   m.GochimpM.Subject = s
   return
}

func (m *MEmail) AddRecipients(r ...string) (err error){
  recipients := []gochimp.Recipient{}
  for _, email := range r {
		recipients = append(recipients, gochimp.Recipient{Email: email})
  }
  m.GochimpM.To=recipients
  return
}

func (m *MEmail) AddAttachment(a string) (err error){
  return
}

func (m *MEmail) AddTemplate(t string)(err error){
	c,err := NewConfig()
    apiKey,err := c.String("mailchimp", "apikey") 
    if err!=nil{
       return err
	}
	mandrillAPI, err := gochimp.NewMandrill(apiKey)
	if err != nil {
		return err
	}

	templateName := "welcome email2"
	contentVar := gochimp.Var{"main", t}
	content := []gochimp.Var{contentVar}

	_, err = mandrillAPI.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		return err
	}
	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, content, nil)
	if err != nil {
		return err
	}
	m.GochimpM.Html= renderedTemplate
    return
}

func (m *MEmail) AddContent(t string)(err error){
    return 
}


func NewConfig() (cfg *config.Config,err error){ 
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