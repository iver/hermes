package providers 

import (
    "errors"
    "github.com/ivan-iver/hermes/providers/mailgun"
    "github.com/ivan-iver/hermes/providers/sendgrid"
    "github.com/ivan-iver/hermes/providers/mailchimp"
)
 
func NewProvider(options []string) (p interface{},err error) { 
  var first = options[0] 
  switch first { 
  case "mailgun": 
     p = mailgun.NewProvider() 
  case "mailchimp": 
     p = mailchimp.NewProvider()
  case "sendgrid": 
     p= sendgrid.NewProvider()
  default: 
     err = errors.New("ERR_UNKNOWN_PROVIDER")
  }
  return  
} 