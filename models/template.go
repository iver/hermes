package models

type Template struct {
	ID   int64   `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

func (t *Template) SetID(id int64){
    t.ID=id
}

func (t *Template) SetName(name string){
   t.Name=name
}

func NewTemplate() *Template{
     return &Template{}
}
