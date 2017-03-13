package sendgrid

 import (
	 mail "github.com/sendgrid/sendgrid-go/helpers/mail"
	 "time"
 )

type SendgridEmail struct{
	ID          int64          `json:"-" db:"id,omitempty"`
	InsertedAt  time.Time      `json:"-" db:"inserted_at,omitempty"`
	SendedAt    time.Time      `json:"-" db:"sended_at,omitempty"`
	SendgridM   *mail.SGMailV3 `json:"-" db:"sendgrid_message,omitempty"`
}


func (m *SendgridEmail) AddSenderEmail(e string) (err error){
   m.SendgridM.From.Address=e
   return
}

func (m *SendgridEmail) AddSenderName(n string) (err error){
   m.SendgridM.From.Name= n
   return
}

func (m *SendgridEmail) AddSubject(s string) (err error){
   m.SendgridM.Subject = s
   return
}

func (m *SendgridEmail) AddRecipients(r ...string) (err error){
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

func (m *SendgridEmail) AddAttachment(a string) (err error){
  return
}

func (m *SendgridEmail) AddTemplate(t string)(err error){
    return
}


func NewEmail() SendgridEmail {
    m:= SendgridEmail{}
	from:=mail.Email{}
	m.SendgridM = mail.NewV3Mail()
	m.SendgridM.SetFrom(&from)
	return m 
}

func (m *SendgridEmail) AddContent(t string)(err error){
	content := mail.NewContent("text/plain", t)
	m.SendgridM.AddContent(content)
    return 
}