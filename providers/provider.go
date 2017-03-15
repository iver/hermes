package providers

func NewProvider(options []string) interface{} {
	var first = options[0]
	switch first {
	case "mailgun":
		// Create an empty mailgun model/provider
	case "mailchimp":
		// Create an empty mailchimp model/provider
	case "sendgrid":
		// Create an empty sengrid model/provider
	default:
		// returns default errors
	}
	return
}
