package mailgun

import (
	//"strings"
	"gopkg.in/mailgun/mailgun-go.v1"
	"github.com/ivan-iver/hermes/models"
)

type Email struct {
	PlainEmail  models.Email     `json:"plain_email,omitempty"`
	MailgunM    *mailgun.Message  `json:"mailgun_m,omitempty"`
}


func (m *Email) AddSender(s interface{}) (err error) {
	sender:=s.(models.Sender)
	m.PlainEmail.Sender = &sender
	return
}

func (m *Email) AddSubject(s string) (err error) {
	m.PlainEmail.Subject = &s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	allrecipient:=r.(models.Recipients)
	m.PlainEmail.Recipients=&allrecipient
	//tos:=strings.Join(allrecipient.To,",")
	return
}

func (m *Email) AddAttachment(a interface{}) (err error) {
	attachment:=a.(models.Attachment)
	m.PlainEmail.Attachments = append(m.PlainEmail.Attachments,&attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template:= t.(models.Template)
	m.PlainEmail.Template=&template
	return
}

func (m *Email) AddContent(c interface{}) (err error) {
	content:= c.(models.Content)
	m.PlainEmail.Content = append(m.PlainEmail.Content,&content)
	return
}

func (m *Email) SetValues() (err error) {
	message1 := mailgun.NewMessage(
		m.PlainEmail.Sender.Email,
		*m.PlainEmail.Subject,
		m.PlainEmail.Content[0].Value,
		"mau.cdr.19@gmail.com")
	m.MailgunM = message1
	return
}

func (m *Email) GetPlainEmail() (email interface{}){
   email=&m.PlainEmail
   return
}
