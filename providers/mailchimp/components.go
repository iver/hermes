package mailchimp


type Attachment struct {
	Type     string  `json:"type,omitempty"`
	Content  string  `json:"content,omitempty"`
	Filename string  `json:"content,omitempty"`
	Filepath string  `json:"content,omitempty"`
}

type Content struct {
	Type  string  `json:"type,omitempty"`
	Value string  `json:"value,omitempty"`
}

type Template struct {
	ID   int64   `json:"id,omitempty"`
	Name string  `json:"id,omitempty"`
}

type Sender struct {
	Name  string  `json:"id,omitempty"`
	Email string  `json:"id,omitempty"`
}

type Recipients struct {
	To   []string  `json:"id,omitempty"`
	CC   []string  `json:"id,omitempty"`
	BCC  []string  `json:"id,omitempty"`
}
	