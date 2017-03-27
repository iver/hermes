package models

import (
	"time"
)

// EEmail struct provide plain email information
type Email struct {
	ID          int64         `json:"id,omitempty"`
	Sender      *Sender       `json:"sender,omitempty"`
	Subject     *string       `json:"subject,omitempty"`
	Content     []*Content    `json:"content,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
	Recipients  *Recipients   `json:"recipients,omitempty"`
	Template    *Template     `json:"template,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	SendedAt    time.Time     `json:"sended_at,omitempty"`
}

