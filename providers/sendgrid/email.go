package sendgrid

import (
	"time"
    "fmt"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	ID          int64          `json:"id,omitempty"`
	Sender      *Sender        `json:"sender,omitempty"`
	Content     []*Content     `json:"content,omitempty"`
	Attachments []*Attachment  `json:"attachments,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	SendedAt    *time.Time     `json:"sended_at,omitempty"`
	SendgridM   *mail.SGMailV3 `json:"id,omitempty"`
}

func (m *Email) AddSender(s interface{}) (err error) {
	sender:=s.(Sender)
	m.SendgridM.From.Name = sender.Name
	m.SendgridM.From.Address = sender.Email
	return
}


func (m *Email) AddSubject(s string) (err error) {
	m.SendgridM.Subject = s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	allrecipient:=r.(Recipients)
	
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
	attachment:=a.(Attachment)
	fmt.Println(attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
	template:=t.(Template)
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
	content:=c.(Content)
	contentM := mail.NewContent("text/plain", content.Value)
	m.SendgridM.AddContent(contentM)
	return
}
