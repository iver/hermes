package models

//interface provider
type Provider interface {
	sendEmail(email Email) error
}
