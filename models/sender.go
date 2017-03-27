package models

// Sender struct provide basic sender information
type Sender struct {
	Name  string  `json:"name,omitempty"`
	Email string  `json:"email,omitempty"`
}

// SetName set sender name
func (s *Sender) SetName(name string) {
   s.Name= name
}

// SetEmail set sender email address
func (s *Sender) SetEmail(email string) {
  s.Email= email
}

// NewSender creates a new sender struct
func NewSender() *Sender{
     return &Sender{}
}
