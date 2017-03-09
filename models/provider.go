package models

// Provider interface
type Provider interface {
	sendEmail(email Email) error
}
