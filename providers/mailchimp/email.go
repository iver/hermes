package mailchimp

import (
	"fmt"
	"github.com/mattbaird/gochimp"
	"github.com/mauricio-cdr/config"
	"github.com/ivan-iver/hermes/models"
)

type Email struct {
    PlainEmail  models.Email    `json:"plain_email,omitempty"`
	GochimpM    gochimp.Message `json:"gochimp_m,omitempty"`
}

func (m *Email) AddSender(s interface{}) (err error) {
	sender,ok := s.(models.Sender)
	if !ok{
		err=models.ErrInvalidSender
		return 
	}
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
	recipient,ok:= r.(models.Recipients)
	if !ok{
		return models.ErrInvalidRecipients
	}
	m.PlainEmail.Recipients= &recipient
	recipients := []gochimp.Recipient{}
	for _, email := range recipient.To {
		recipients = append(recipients, gochimp.Recipient{Email: email})
	}
	m.GochimpM.To = recipients
	return
}

func (m *Email) AddAttachment(a interface{}) (err error) {
	attachment,ok:= a.(models.Attachment)
	if !ok{
		return models.ErrInvalidAttachment
	}
	m.PlainEmail.Attachments = append(m.PlainEmail.Attachments,&attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template,ok:= t.(models.Template)
	if !ok{
		return models.ErrInvalidTemplate
	}
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
	content,ok := c.(models.Content)
	if !ok{
		return models.ErrInvalidContent
	}
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

func (e *Email) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"id":          e.PlainEmail.ID,
		"sender":      e.PlainEmail.Sender,
		"subject":     e.PlainEmail.Subject,
		"content":     e.PlainEmail.Content,
		"attachments": e.PlainEmail.Attachments,
		"recipients":  e.PlainEmail.Recipients,
		"template":    e.PlainEmail.Template,
		"created_at":  e.PlainEmail.CreatedAt,
		"sended_at":   e.PlainEmail.SendedAt,
	}
}
