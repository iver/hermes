package models

type Recipients struct {
	To   []string  `json:"to,omitempty"`
	CC   []string  `json:"cc,omitempty"`
	BCC  []string  `json:"bcc,omitempty"`
}

func (r *Recipients) AddTos(to ...string) {
	r.To = append(r.To, to...)
}

func (r *Recipients) AddCCs(cc ...string) {
	r.CC = append(r.CC, cc...)
}

func (r *Recipients) AddBCCs(bcc ...string) {
	r.BCC = append(r.BCC, bcc...)
}

func NewRecipients() *Recipients{
     return &Recipients{
         To:            make([]string, 0),
		 CC:            make([]string, 0),
		 BCC:           make([]string, 0),
	 }
}
