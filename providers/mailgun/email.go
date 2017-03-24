package mailgun

import (
	"fmt"
	"time"
	"strings"
	"gopkg.in/mailgun/mailgun-go.v1"
	"github.com/ivan-iver/hermes/models"
)

type Email struct {
	ID          *int64               `json:"id,omitempty"`
	Sender      *models.Sender       `json:"sender,omitempty"`
	Subject     *string              `json:"dubject,omitempty"`
	Attachments []*models.Attachment `json:"attachments,omitempty"`
	Template    *models.Template     `json:"template,omitempty"`
	Content     *models.Content      `json:"content,omitempty"`
	InsertedAt  time.Time            `json:"inserted_at,omitempty"`
	SendedAt    *time.Time           `json:"sended_at,omitempty"`
	MailgunM    *mailgun.Message     `json:"mailgun_m,omitempty"`
}


func (m *Email) AddSender(s interface{}) (err error) {
	sender:=s.(models.Sender)
	m.Sender = &sender
	return
}

func (m *Email) AddSubject(s string) (err error) {
	m.Subject = &s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	allrecipient:=r.(models.Recipients)
	tos:=strings.Join(allrecipient.To,",")
	fmt.Println(tos)
	return
}

func (m *Email) AddAttachment(a interface{}) (err error) {
	attachment:=a.(models.Attachment)
	fmt.Println(attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template:= t.(models.Template)
	fmt.Println(template)
	return
}

func (m *Email) AddContent(c interface{}) (err error) {
	content:= c.(models.Content)
	m.Content= &content
	return
}

func (m *Email) SetValues() (err error) {
	message1 := mailgun.NewMessage(
		m.Sender.Email,
		*m.Subject,
		m.Content.Value,
		"mau.cdr.19@gmail.com")
	m.MailgunM = message1
	return
}
