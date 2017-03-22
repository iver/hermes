package mailgun

 import (
	 "time"
	 "gopkg.in/mailgun/mailgun-go.v1"
 )

type Email struct{
	ID          *int64            `json:"-" db:"-,omitempty"`
	senderName  *string           `json:"-" db:"-,omitempty"`
	senderEmail *string           `json:"-" db:"-,omitempty"`
	subject     *string           `json:"-" db:"-,omitempty"`
    attachment  *[]string         `json:"-" db:"-,omitempty"`
    template    *string           `json:"-" db:"-,omitempty"`
	content     *string           `json:"-" db:"-,omitempty"`
	InsertedAt  time.Time         `json:"-" db:"inserted_at,omitempty"`
	SendedAt    *time.Time        `json:"-" db:"sended_at,omitempty"`
	MailgunM    *mailgun.Message  `json:"-" db:"mailgun_message,omitempty"`
}


func (m *Email) AddSenderEmail(e string) (err error){
   m.senderEmail=&e	
   return
}

func (m *Email) AddSenderName(n string) (err error){
    m.senderName=&n	
   return
}

func (m *Email) AddSubject(s string) (err error){
    m.subject=&s
   return
}

func (m *Email) AddRecipients(r ...string) (err error){
    /*recipients:="mau.cdr.19@gmail.com"
	m.MailgunM.AddRecipient(recipients)*/
	return
}

func (m *Email) AddAttachment(a string) (err error){
  return
}

func (m *Email) AddTemplate(t string)(err error){
    m.template=&t
	return
}

func (m *Email) AddContent(t string)(err error){
	m.content=&t
    return 
}

func (m *Email) SetValues()(err error){
	message1 := mailgun.NewMessage(
		*m.senderEmail,
		*m.subject,
		*m.content,
		"mau.cdr.19@gmail.com")	
	m.MailgunM=message1	
    return 
}