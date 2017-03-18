package lib

// Email interface
type Email interface {
   AddSenderEmail(e string) error
   AddSubject(s string) error
   AddSenderName(name string) error
   AddRecipients(e ...string) error
   AddAttachment(p string) error
   AddTemplate(t string) error
   AddContent(c string) error
}
