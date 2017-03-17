package sendgrid

 import (
	 mail "github.com/sendgrid/sendgrid-go/helpers/mail"
	 "time"
 )

type Email struct{
	ID          int64          `json:"-" db:"id,omitempty"`
	InsertedAt  time.Time      `json:"-" db:"inserted_at,omitempty"`
	SendedAt    time.Time      `json:"-" db:"sended_at,omitempty"`
	SendgridM   *mail.SGMailV3 `json:"-" db:"sendgrid_message,omitempty"`
}


func (m *Email) AddSenderEmail(e string) (err error){
   m.SendgridM.From.Address=e
   return
}

func (m *Email) AddSenderName(n string) (err error){
   m.SendgridM.From.Name= n
   return
}

func (m *Email) AddSubject(s string) (err error){
   m.SendgridM.Subject = s
   return
}

func (m *Email) AddRecipients(r ...string) (err error){
    recipients := []*mail.Email{}

	for _, email := range r {
        eml := mail.Email{Address: email}
		recipients = append(recipients, &eml )
	}
	ps := mail.NewPersonalization()
    ps.AddTos(recipients...)
    m.SendgridM.AddPersonalizations(ps)
	return
}

func (m *Email) AddAttachment(a string) (err error){
  return
}

func (m *Email) AddTemplate(t string)(err error){
    return
}


func NewEmail() Email {
    m:= Email{}
	from:=mail.Email{}
	m.SendgridM = mail.NewV3Mail()
	m.SendgridM.SetFrom(&from)
	return m 
}

func (m *Email) AddContent(t string)(err error){
	content := mail.NewContent("text/plain", t)
	m.SendgridM.AddContent(content)
    return 
}