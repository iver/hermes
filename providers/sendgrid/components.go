package sendgrid


type Attachment struct {
	Type     string  `json:"type,omitempty"`
	Content  string  `json:"content,omitempty"`
	Filename string  `json:"filename,omitempty"`
	Filepath string  `json:"filepath,omitempty"`
}

type Content struct {
	Type  string  `json:"type,omitempty"`
	Value string  `json:"value,omitempty"`
}

type Template struct {
	ID   int64   `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

type Sender struct {
	Name  string  `json:"name,omitempty"`
	Email string  `json:"email,omitempty"`
}

type Recipients struct {
	To   []string  `json:"to,omitempty"`
	CC   []string  `json:"cc,omitempty"`
	BCC  []string  `json:"bcc,omitempty"`
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

func (c *Content) SetType(typec string){
   c.Type=typec
}

func (c *Content) SetValue(value string){
   c.Value=value
}

func (t *Template) SetID(id int64){
    t.ID=id
}

func (t *Template) SetName(name string){
   t.Name=name
}

func (s *Sender) SetName(name string) {
   s.Name= name
}

func (s *Sender) SetEmail(email string) {
  s.Email= email
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

func NewAttachment() *Attachment{
    return &Attachment{}
}

func NewContent() *Content{
    return &Content{}
}

func NewTemplate() *Template{
     return &Template{}
}

func NewSender() *Sender{
     return &Sender{}
}

func NewRecipients() *Recipients{
     return &Recipients{
         To:            make([]string, 0),
		 CC:            make([]string, 0),
		 BCC:           make([]string, 0),
	 }
}