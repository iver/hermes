package models

// Attachment struct provide basic attachment information
type Attachment struct {
	Type     string  `json:"type,omitempty"`
	Content  string  `json:"content,omitempty"`
	Filename string  `json:"filename,omitempty"`
	Filepath string  `json:"filepath,omitempty"`
}

// SetType set attachment type
func (a *Attachment) SetType(typeA string){
   a.Type=typeA
}

// SetContent set a information about attachment
func (a *Attachment) SetContent(content string){
   a.Content=content
}

// SetFilename set a name for attachment
func (a *Attachment) SetFilename(filename string){
   a.Filename=filename
}

// SetFilepath set the attachment location path
func (a *Attachment) SetFilepath(filepath string){
   a.Filepath=filepath
}

// NewAttachment creates a new attachment struct
func NewAttachment() *Attachment{
    return &Attachment{}
}