package models

type Email struct {
   SenderName   string
   SenderEmail  string
   Recipients   []string
   Subject      string
   Content      string
}