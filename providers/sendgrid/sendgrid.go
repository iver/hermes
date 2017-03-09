package providers

import (
	"os"
		"fmt"
		"log"
	"github.com/ivan-iver/hermes/models"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go"
)

type Sendgrid struct {
}

// SendEmail method with mailgun provider
func (p *Sendgrid) SendEmail(email models.Email) (er error) {

	APIKey := os.Getenv("SENDGRID_API_KEY")
	from := mail.NewEmail(email.SenderName, email.SenderEmail)
	subject := email.Subject
	content := mail.NewContent("text/plain", email.Template)

	recipients := []*mail.Email{}

	for _, email := range email.Recipients {
        eml := mail.Email{Address: email}
		recipients = append(recipients, &eml )
	}

	m := mail.NewV3Mail()
	m.SetFrom(from)
	m.Subject = subject
	ps := mail.NewPersonalization()
    ps.AddTos(recipients...)
	m.AddPersonalizations(ps)
	m.AddContent(content)
	request := sendgrid.GetRequest(APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
        log.Println(err)
		return err
    } else {
        fmt.Println(response.StatusCode)
        fmt.Println(response.Body)
        fmt.Println(response.Headers)
    }
	return
}
