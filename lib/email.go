package lib

// Email interface
type Email interface {
   AddSender(s interface{}) error
   AddSubject(s string) error
   AddRecipients(r interface{}) error
   AddAttachment(a interface{}) error
   AddTemplate(t interface{}) error
   AddContent(c interface{}) error
}
