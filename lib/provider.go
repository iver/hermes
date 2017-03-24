package lib

// Provider interface
type Provider interface {
	Init() error
	GetName() string 
	SendEmail(email interface{}) error
	NewEmail(sn interface{}, s string ,t interface{}) (email interface{},err error)
	ToString() string
}