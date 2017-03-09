package providers

import (
	"fmt"
    "os"
	"github.com/ivan-iver/hermes/models"
	"github.com/mattbaird/gochimp"
)

type Mailchimp struct {
}

//  sendemail function with mailchimp provider
func (p *Mailchimp) SendEmail(email models.Email) (err error) {

    apiKey := os.Getenv("MANDRILL_KEY")
	mandrillAPI, err := gochimp.NewMandrill(apiKey)
	if err != nil {
		fmt.Println("Error instantiating client")
		return err
	}

	templateName := "welcome email"
	contentVar := gochimp.Var{"main", email.Content}
	content := []gochimp.Var{contentVar}

	_, err = mandrillAPI.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		fmt.Printf("Error adding template: %v", err)
		return err
	}

	defer mandrillAPI.TemplateDelete(templateName)

	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, content, nil)
	if err != nil {
		fmt.Printf("Error rendering template: %v", err)
		return err
	}
	eml := email.Recipients[0]
	recipients := []gochimp.Recipient{
		gochimp.Recipient{Email: eml},
	}

	message := gochimp.Message{
		Html:      renderedTemplate,
		Subject:   email.Subject,
		FromEmail: email.SenderEmail,
		FromName:  email.SenderName,
		To:        recipients,
	}

	if _, err = mandrillAPI.MessageSend(message, false); err != nil {
		fmt.Println("Error sending message")
		return err
	}

	return
}
