package server

import "github.com/hiroaki-yamamoto/reusable-services/email/sendfuncs"

// Server represents a service of sending important emails.
type Server struct {
	// SendFunction is a function to send an email.
	// If sending an email is failed, the next function is invoke.
	// i.e. func[0] -failed-> func[1] -failed-> func[2] -failed->...func[n]
	SendFuncs []sendfuncs.Send
	From      string
}

// New creates a new instance of Email Service.
func New(from string, sendfuncs ...sendfuncs.Send) *Server {
	return &Server{
		SendFuncs: sendfuncs,
		From:      from,
	}
}
