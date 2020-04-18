package server

import "github.com/hiroaki-yamamoto/reusable-services/email/sendfuncs"

// Email represents a service of sending important emails.
type Email struct {
	// SendFunction is a function to send an email.
	// If sending an email is failed, the next function is invoke.
	// i.e. func[0] -failed-> func[1] -failed-> func[2] -failed->...func[n]
	SendFuncs []sendfuncs.Send
}

// New creates a new instance of Email Service.
func New(sendfuncs ...sendfuncs.Send) *Email {
	return &Email{SendFuncs: sendfuncs}
}
