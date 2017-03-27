package models

// Content struct provide all recipients email information
type Recipients struct {
	To   []string  `json:"to,omitempty"`
	CC   []string  `json:"cc,omitempty"`
	BCC  []string  `json:"bcc,omitempty"`
}

// AddTos struct Add a email list of common recipients
func (r *Recipients) AddTos(to ...string) {
	r.To = append(r.To, to...)
}

// AddTos struct Add a email list of CC recipients
func (r *Recipients) AddCCs(cc ...string) {
	r.CC = append(r.CC, cc...)
}

// AddTos struct Add a email list of BCC recipients
func (r *Recipients) AddBCCs(bcc ...string) {
	r.BCC = append(r.BCC, bcc...)
}

//NewRecipients creates a new recipients struct with empty list
func NewRecipients() *Recipients{
     return &Recipients{
         To:            make([]string, 0),
		 CC:            make([]string, 0),
		 BCC:           make([]string, 0),
	 }
}
