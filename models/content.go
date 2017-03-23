package models

type Content struct {
	Type  string  `json:"type,omitempty"`
	Value string  `json:"value,omitempty"`
}

func (c *Content) SetType(typec string){
   c.Type=typec
}

func (c *Content) SetValue(value string){
   c.Value=value
}

func NewContent() *Content{
    return &Content{}
}