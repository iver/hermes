package models


// Email interface
type Email interface {
   AddSenderEmail(e String) error
   AddSubject(s string) error
   AddSenderName(name string) error
   AddRecipients(e ..string) error
   AddAttachment(path string) error
   AddTemplate(t string) error
   AddContent(c string) error
}
