package mailchimp

import (
	"fmt"
	"github.com/mattbaird/gochimp"
	"github.com/mauricio-cdr/config"
	"github.com/ivan-iver/hermes/models"
)

type Email struct {
    PlainEmail  models.Email    `json:"plain_email,omitempty"`
	GochimpM    gochimp.Message  `json:"gochimp_m,omitempty"`
}

func (m *Email) AddSender(s interface{}) (err error) {
	sender := s.(models.Sender)
	m.GochimpM.FromEmail = sender.Email
	m.GochimpM.FromName = sender.Name
	m.PlainEmail.Sender =&sender
	return
}

func (m *Email) AddSubject(s string) (err error) {
	m.PlainEmail.Subject = &s
	m.GochimpM.Subject = s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	recipient := r.(models.Recipients)
	m.PlainEmail.Recipients= &recipient
	recipients := []gochimp.Recipient{}
	for _, email := range recipient.To {
		recipients = append(recipients, gochimp.Recipient{Email: email})
	}
	m.GochimpM.To = recipients
	return
}

func (m *Email) AddAttachment(a interface{}) (err error) {
	attachment := a.(models.Attachment)
	m.PlainEmail.Attachments = append(m.PlainEmail.Attachments,&attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template:= t.(models.Template)
	m.PlainEmail.Template=&template

	c, err := config.NewConfig()
	apiKey, err := c.Property("mailchimp", "apikey")
	if err != nil {
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
	m.GochimpM.Html = renderedTemplate
	return
}

func (m *Email) AddContent(c interface{}) (err error) {
	content := c.(models.Content)
	m.PlainEmail.Content = append(m.PlainEmail.Content,&content)
	conf, err := config.NewConfig()
	apiKey, err := conf.Property("mailchimp", "apikey")
	if err != nil {
		return err
	}
	mandrillAPI, err := gochimp.NewMandrill(apiKey)
	if err != nil {
		return err
	}

	templateName := "welcome email2"
	contentVar := gochimp.Var{"main", content.Value}
	contentM := []gochimp.Var{contentVar}

	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, contentM, nil)
	if err != nil {
		return err
	}
	m.GochimpM.Html = renderedTemplate
	return
}

func (m *Email) GetPlainEmail() (email interface{}){
   	email=&m.PlainEmail
   return
}

 func (m *Email)RefactorEmail(e interface{}){
	/* pemail=e.(model.PlainEmail)
	 m.PlainEmail=pemail*/
	 return
 }