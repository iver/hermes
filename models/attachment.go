package models

type Attachment struct {
	Type     string  `json:"type,omitempty"`
	Content  string  `json:"content,omitempty"`
	Filename string  `json:"filename,omitempty"`
	Filepath string  `json:"filepath,omitempty"`
}

func (a *Attachment) SetType(typeA string){
   a.Type=typeA
}

func (a *Attachment) SetContent(content string){
   a.Content=content
}

func (a *Attachment) SetFilename(filename string){
   a.Filename=filename
}

func (a *Attachment) SetFilepath(filepath string){
   a.Filepath=filepath
}

func NewAttachment() *Attachment{
    return &Attachment{}
}