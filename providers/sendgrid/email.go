package sendgrid

import (
	"github.com/iver/hermes/models"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	PlainEmail models.Email   `json:"plain_email,omitempty"`
	SendgridM  *mail.SGMailV3 `json:"id,omitempty"`
}

func (m *Email) AddSender(s interface{}) (err error) {
	sender, ok := s.(models.Sender)
	if !ok {
		return models.ErrInvalidSender
	}
	m.PlainEmail.Sender = &sender
	m.SendgridM.From.Name = sender.Name

	m.SendgridM.From.Address = sender.Email
	return
}

func (m *Email) AddSubject(s string) (err error) {
	m.SendgridM.Subject = s
	m.PlainEmail.Subject = &s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	allrecipient, ok := r.(models.Recipients)
	if !ok {
		return models.ErrInvalidRecipients
	}
	m.PlainEmail.Recipients = &allrecipient
	recipients := []*mail.Email{}

	for _, email := range allrecipient.To {
		eml := mail.Email{Address: email}
		recipients = append(recipients, &eml)
	}
	ps := mail.NewPersonalization()
	ps.AddTos(recipients...)
	m.SendgridM.AddPersonalizations(ps)
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

func NewEmail() Email {
	m := Email{}
	from := mail.Email{}
	m.SendgridM = mail.NewV3Mail()
	m.SendgridM.SetFrom(&from)
	return m
}

func (m *Email) AddContent(c interface{}) (err error) {
	content, ok := c.(models.Content)
	if !ok {
		return models.ErrInvalidContent
	}
	m.PlainEmail.Content = append(m.PlainEmail.Content, &content)
	contentM := mail.NewContent("text/plain", content.Value)
	m.SendgridM.AddContent(contentM)
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
