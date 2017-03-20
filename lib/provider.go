package lib

// Provider interface
type Provider interface {
	Init() error
	GetName() string 
	SendEmail(email interface{}) error
	NewEmail(se string , sn string , s string ,t string) (email interface{},err error)
	ToString() string
}