package models

type Template struct {
	ID   int64   `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

// SetID set sender id
func (t *Template) SetID(id int64){
    t.ID=id
}

// SetName set sender name
func (t *Template) SetName(name string){
   t.Name=name
}

// NewTemplate creates a new sender struct
func NewTemplate() *Template{
     return &Template{}
}
