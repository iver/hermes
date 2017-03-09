package providers

import (
	"os"
	"github.com/ivan-iver/hermes/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Sendgrid struct {
}

// SendEmail method with mailgun provider
func (p *Sendgrid) SendEmail(email models.Email) (err error) {

	APIKey :=   os.Getenv("SENDGRID_API_KEY")

	from := mail.NewEmail(email.SenderName, email.SenderEmail)
	subject := email.Subject
	to := mail.NewEmail("You", email.Recipients[0])
	content := mail.NewContent("text/plain", email.Content)
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err = sendgrid.API(request)
	if err != nil {
		return err
	}

	return
}
