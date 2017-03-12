package mailchimp
 import (
	 "github.com/mattbaird/gochimp"
	 "time"
	 "os"
	 "fmt"
 )

type MailchimpEmail struct{
	ID          int64          `json:"-" db:"id,omitempty"`
	InsertedAt  time.Time      `json:"-" db:"inserted_at,omitempty"`
	SendedAt    time.Time      `json:"-" db:"sended_at,omitempty"`
	GochimpM    gochimp.Message`json:"-" db:"gochimp_message,omitempty"`
}


func (m *MailchimpEmail) AddSenderEmail(e string) (err error){
   m.GochimpM.FromEmail=e
   return
}

func (m *MailchimpEmail) AddSenderName(n string) (err error){
   m.GochimpM.FromName= n
   return
}

func (m *MailchimpEmail) AddSubject(s string) (err error){
   m.GochimpM.Subject = s
   return
}

func (m *MailchimpEmail) AddRecipients(r ...string) (err error){
  recipients := []gochimp.Recipient{}
  for _, email := range r {
		recipients = append(recipients, gochimp.Recipient{Email: email})
  }
  m.GochimpM.To=recipients
  return
}

func (m *MailchimpEmail) AddAttachment(a string) (err error){
  return
}

func (m *MailchimpEmail) AddTemplate(t string)(err error){
    apiKey := os.Getenv("MANDRILL_KEY")
	mandrillAPI, err := gochimp.NewMandrill(apiKey)
	if err != nil {
		return err
	}

	templateName := "welcome email"
	contentVar := gochimp.Var{"main", t}
	content := []gochimp.Var{contentVar}

	_, err = mandrillAPI.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		fmt.Printf("Error adding template: %v", err)
		return err
	}
	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, content, nil)
	if err != nil {
		fmt.Printf("Error rendering template: %v", err)
		return err
	}
	m.GochimpM.Html= renderedTemplate
    return
}

func NewEmail() MailchimpEmail {
	return 
}

func (m *MailchimpEmail) AddContent(t string)(err error){
    return 
}