package models

// Provider interface
type Provider interface {
	Init() error
	SendEmail(email Email) error
	NewEmail(se string , sn string , s string ,t string) (m Email,err error)
