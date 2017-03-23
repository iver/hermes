package models

type Sender struct {
	Name  string  `json:"name,omitempty"`
	Email string  `json:"email,omitempty"`
}

func (s *Sender) SetName(name string) {
   s.Name= name
}

func (s *Sender) SetEmail(email string) {
  s.Email= email
}

func NewSender() *Sender{
     return &Sender{}
}
