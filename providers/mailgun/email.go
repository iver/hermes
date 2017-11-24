package mailgun

import (
	"github.com/iver/hermes/models"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type Email struct {
	PlainEmail models.Email     `json:"plain_email,omitempty"`
	MailgunM   *mailgun.Message `json:"mailgun_m,omitempty"`
}

func (m *Email) AddSender(s interface{}) (err error) {
	sender, ok := s.(models.Sender)
	if !ok {
		return models.ErrInvalidSender
	}
	m.PlainEmail.Sender = &sender
	return
}

func (m *Email) AddSubject(s string) (err error) {
	m.PlainEmail.Subject = &s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	allrecipient, ok := r.(models.Recipients)
	if !ok {
		return models.ErrInvalidRecipients
	}
	m.PlainEmail.Recipients = &allrecipient
	return
}

func (m *Email) AddAttachment(a interface{}) (err error) {
	attachment, ok := a.(models.Attachment)
	if !ok {
		return models.ErrInvalidAttachment
	}
	m.PlainEmail.Attachments = append(m.PlainEmail.Attachments, &attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template, ok := t.(models.Template)
	if !ok {
		return models.ErrInvalidTemplate
	}
	m.PlainEmail.Template = &template
	return
}

func (m *Email) AddContent(c interface{}) (err error) {
	content, ok := c.(models.Content)
	if !ok {
		return models.ErrInvalidContent
	}
	m.PlainEmail.Content = append(m.PlainEmail.Content, &content)
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

func (m *Email) GetPlainEmail() (email interface{}) {
	email = &m.PlainEmail
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
