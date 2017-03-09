package providers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ivan-iver/hermes/models"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type Mailgun struct {
}

// send email method with mailgun provider

func (p *Mailgun) SendEmail(email models.Email) (err error) {

	domain := os.Getenv("MG_DOMAIN")
	APIKey := os.Getenv("MG_API_KEY")
	publicAPIKey := os.Getenv("MG_PUBLIC_API_KEY")

	mg := mailgun.NewMailgun(domain, APIKey, publicAPIKey)
	message := mailgun.NewMessage(
		email.SenderEmail,
		email.Subject,
		"content",
		strings.Join(email.Recipients, ","))

	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	return
}
