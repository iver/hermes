package mailchimp

import (
	"fmt"
	"time"
	"github.com/mattbaird/gochimp"
	"github.com/mauricio-cdr/config"
	"github.com/ivan-iver/hermes/models"
)

type Email struct {
	ID         int64           `json:"-" db:"id,omitempty"`
	InsertedAt time.Time       `json:"-" db:"inserted_at,omitempty"`
	SendedAt   time.Time       `json:"-" db:"sended_at,omitempty"`
	GochimpM   gochimp.Message `json:"-" db:"gochimp_message,omitempty"`
}


func (m *Email) AddSender(s interface{}) (err error) {
	sender:=s.(models.Sender)
	m.GochimpM.FromEmail = sender.Email
	m.GochimpM.FromName  = sender.Name
	return
}

func (m *Email) AddSubject(s string) (err error) {
	m.GochimpM.Subject = s
	return
}

func (m *Email) AddRecipients(r interface{}) (err error) {
	recipient:= r.(models.Recipients)
	recipients := []gochimp.Recipient{}
	for _, email := range recipient.To {
		recipients = append(recipients, gochimp.Recipient{Email: email})
	}
	m.GochimpM.To = recipients
	return
}

func (m *Email) AddAttachment(a interface{}) (err error) {
	attachment:=a.(models.Attachment)
	fmt.Println(attachment)
	return
}

func (m *Email) AddTemplate(t interface{}) (err error) {
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
	content:=c.(models.Content)
	fmt.Println(content)
	return
}
