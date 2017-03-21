package hermes_test


type Email struct{
   	SenderEmail  string
	SenderName   string
	Subject      string
	Content      string
	Recipients   []string
}

//Peticiónn para registro de usuario 
func CorrectEmail() (e *Email) {
	e = &Email{
		SenderEmail: "hermes@hermes.com",
		SenderName: "A friend",
		Subject: "Test with hermes",
		Recipients:[]string{"mau.cdr.19@gmail.com","mau_dsx2@hotmail.com"},
	}
	return
}
