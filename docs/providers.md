Frameworks para el envio de email en Go
---

# mailjet

* **Configuración**

```
  export MJ_APIKEY_PUBLIC='your api key'
  export MJ_APIKEY_PRIVATE='your api secret'
```

* **Instalación**
```
     (go get github.com/mailjet/mailjet-apiv3-go)
```

* **Ejemplo de uso**

```
    import (
      "github.com/mailjet/mailjet-apiv3-go"
      "github.com/mailjet/mailjet-apiv3-go/resources"
      "fmt"
      "os"
    )

    func main() {
       publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
       secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

       mj := mailjet.NewMailjetClient(publicKey, secretKey)

      param := &mailjet.InfoSendMail{
          FromEmail: "qwe@qwe.com",
          FromName: "Bob Patrick",
          Recipients: []mailjet.Recipient{
            mailjet.Recipient{
                Email: "qwe@qwe.com",
            },
        },
        Subject: "Hello World!",
        TextPart: "Hi there !",
    }
    res, err := mj.SendMail(param)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Success")
        fmt.Println(res)
    }
}

```

## Posibles respuestas HTTP

| Status | Description |
| ------ | ----------- |
|  200  | OK	All went well. Congrats! |
|  201  |	Created	The POST request was successfully executed. |
|  204	 | No Content	No content found or expected to return. Returned when DELETE was successful.|
|  304	 | Not Modified	The PUT request didn’t affect any record.|
|  400  | Bad Request	One or more parameters are missing or maybe mispelled (unknown resource or action) |
|  401	 | Unauthorized	You have specified an incorrect Api Key / API Secret Key. You may be unauthorized to access the API or your API key may be expired. Visit API  keys Management section to check your keys. |
|  403 |	Forbidden	You are not authorised to access this resource. |
|  404	| Not Found	The resource with the specified ID you are trying to reach does not exist. |
|  405 |	Method Not Allowed	The method requested on the resource does not exist.|
|  429	| Too Many Requests	Oops! You have reach the maximum number of calls allowed per minute by our API. Please review your integration to reduce the number of call issued by your system.|
|  500	| Internal Server Error	Ouch! Something went wrong on our side and we apologize! Please contact our support team who’ll be able to help you on this. |

**IMPORTANTE** (429) es para saber cuando se acaben los mensajes.


# mailgun-go

* **Instalación**

```
 $ go get gopkg.in/mailgun/mailgun-go.v1
```

* **Ejemplo de uso**

```
     import "go get gopkg.in/mailgun/mailgun-go.v1"

     mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
     message := mailgun.NewMessage("sender@example.com", "Fancy subject!", "Hello from Mailgun Go!", "recipient@example.com")
```

* **Respuestas posibles**
 
|  Code	| Description  |
| ------- | ------------ |
|  200	| Everything worked as expected |
|  400		| Bad Request - Often missing a required parameter |
|  401		| Unauthorized - No valid API key provided |
|  402		| Request Failed - Parameters were valid but request failed |
|  404		| Not Found - The requested item doesn’t exist |
|  50(0-4)| Server Errors - something is wrong on Mailgun’s end |
 
 **IMPORTANTE** (402) para saber cuando se acaben los mensajes.
 
 
# sendgrid-go 

* **Configurar el entorno**
   
```
   echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
   echo "sendgrid.env" >> .gitignore
   source ./sendgrid.env
```

* **Instalación**
```
       ( go get github.com/sendgrid/sendgrid-go )  
```

* **Ejemplo de uso** 

```
    import (
    "fmt"
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
    "os"
    )

    func main() {
      from := mail.NewEmail("Example User", "test@example.com")
      subject := "Hello World from the SendGrid Go Library"
      to := mail.NewEmail("Example User", "test@example.com")
      content := mail.NewContent("text/plain", "some text here")
      m := mail.NewV3MailInit(from, subject, to, content)

      request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
      request.Method = "POST"
      request.Body = mail.GetRequestBody(m)
      response, err := sendgrid.API(request)
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Println(response.StatusCode)
        fmt.Println(response.Body)
        fmt.Println(response.Headers)
     }
  }
  
```

* **Posibles respuestas HTTP**

|  Code	| Reason	Description  |
| ------- | ------------------ |
|  200	 | Everything worked as expected |
|  4xx | 4xx responses indicate an error with your request	There was a problem with your request. |
|  400 |	BAD REQUEST	|
|  401 |	UNAUTHORIZED	You do not have authorization to make the request. |
|  403	| FORBIDDEN	|
|  404	| NOT FOUND	The resource you tried to locate could not be found or does not exist. |
|  405	| METHOD NOT ALLOWED	|
|  413 |	PAYLOAD TOO LARGE	The JSON payload you have included in your request is too large.|
|  429	| TOO MANY REQUESTS	The number of requests you have made exceeds SendGrid’s rate limitations. |
|  5xx |	5xx responses indicate an error made by SendGrid	The request you made was valid, but an error occurred when SendGrid attempted to process it.|
|  500 |	SERVER ERROR	An error occurred on a SendGrid server. |
|  503 |	SERVICE NOT AVAILABLE	The SendGrid v3 Web API is not available. |

**IMPORTANTE** (429) para saber cuando se acaben los mensajes.

# gochimp

* **Configuración**

```
   export MANDRILL_KEY=111111111-1111-1111-1111-111111111
   export MANDRILL_USER=user@domain.com
```

* **Ejemplo de uso**

```
   import (
     "fmt"
      "github.com/mattbaird/gochimp"
      "os"
   )

   func main() {
      apiKey := os.Getenv("MANDRILL_KEY")
      mandrillApi, err := gochimp.NewMandrill(apiKey)
 
      if err != nil {
        fmt.Println("Error instantiating client")
      }

      templateName := "welcome email"
      contentVar := gochimp.Var{"main", "<h1>Welcome aboard!</h1>"}
      content := []gochimp.Var{contentVar}

      _, err = mandrillApi.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
      if err != nil {
        fmt.Println("Error adding template: %v", err)
        return
      }
      defer mandrillApi.TemplateDelete(templateName)
      renderedTemplate, err := mandrillApi.TemplateRender(templateName, content, nil)

      if err != nil {
        fmt.Println("Error rendering template: %v", err)
        return
      }

      recipients := []gochimp.Recipient{
        gochimp.Recipient{Email: "person@place.com"},
      }

      message := gochimp.Message{
        Html:      renderedTemplate,
        Subject:   "Welcome aboard!",
        FromEmail: "person@place.com",
        FromName:  "Boss Man",
        To:        recipients,
      }

      _, err = mandrillApi.MessageSend(message, false)

      if err != nil {
        fmt.Println("Error sending message")
      }
}
```


### Recursos de frameworks

* [mailjet](https://github.com/mailjet/mailjet-apiv3-go)
* [mailgun](https://github.com/mailgun/mailgun-go)
* [sendgrid](https://github.com/sendgrid/sendgrid-go)
* [mailchimp](https://github.com/mattbaird/gochimp)

### Recursos API Proveedores

* [maijet](https://dev.mailjet.com/guides/#about-the-mailjet-restful-api)
* [mailgun](https://documentation.mailgun.com/)
* [sengrid](https://sendgrid.com/docs/API_Reference/index.html)

* **Funcionamiento esperado**

1. Enviar los mails a travéz de de mailjet hasta tener un error 429
2. Cambiar para usar mailgun y continuar enviando 
3. Enviar los mails a travéz de mailgun hasta tener un error 402
4. Cambiar para usar sendgrid y continuar enviando   hasta obtener un error 429











