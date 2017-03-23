package sendgrid

import (
    "fmt"
	"time"
	"github.com/ivan-iver/hermes/models"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	ID          int64                `json:"id,omitempty"`
	Sender      *models.Sender       `json:"sender,omitempty"`
	Content     []*models.Content    `json:"content,omitempty"`
	Attachments []*models.Attachment `json:"attachments,omitempty"`
	CreatedAt   *time.Time           `json:"created_at,omitempty"`
	SendedAt    *time.Time           `json:"sended_at,omitempty"`
	SendgridM   *mail.SGMailV3       `json:"id,omitempty"`
}

func (m *Email) AddSender(s interface{}) (err error) {
	sender:=s.(models.Sender)
	m.SendgridM.From.Name = sender.Name
	m.SendgridM.From.Address = sender.Email
	return
}


func (m *Email) AddSubject(s string) (err error) {
	m.SendgridM.Subject = s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	allrecipient:=r.(models.Recipients)
	
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
	attachment:=a.(models.Attachment)
	fmt.Println(attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template:=t.(models.Template)
	fmt.Println(template)
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
	content:=c.(models.Content)
	contentM := mail.NewContent("text/plain", content.Value)
	m.SendgridM.AddContent(contentM)
	return
}
