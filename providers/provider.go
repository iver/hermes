package providers 

import (
    "errors"
    "github.com/ivan-iver/hermes/lib"
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

func Equals(providerI1 interface{},providerI2 interface{})bool{
    var ok bool
    var provider1 lib.Provider
    var provider2 lib.Provider
    if provider1,ok = providerI1.(lib.Provider); !ok {
       return false
    }

    if provider2,ok = providerI2.(lib.Provider);!ok{
       return false
    }

    return (provider1.ToString() == provider2.ToString())

}