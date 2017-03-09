package models

// Email structure
type Email struct {
	SenderName  string
	SenderEmail string
	Recipients  []string
	Subject     string
	Content     string
	Attachments interface{}
	Template    string
}
