package models

// Content struct provide basic content information
type Content struct {
	Type  string  `json:"type,omitempty"`
	Value string  `json:"value,omitempty"`
}

// SetType set content type
func (c *Content) SetType(typec string){
   c.Type=typec
}

// SetValue set the body content 
func (c *Content) SetValue(value string){
   c.Value=value
}

// NewContent creates a new content struct
func NewContent() *Content{
    return &Content{}
}