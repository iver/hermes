package models

// Provider interface
type Provider interface {
	Init() error
	GetName() string 
	SendEmail(email Email) error
	NewEmail(se string , sn string , s string ,t string) (err error)

}