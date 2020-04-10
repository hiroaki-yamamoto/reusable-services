package server

import "github.com/hiroaki-yamamoto/reusable-services/auth/crypto"

// TemplateMap represents a structure that has template names corresponding
// to the templates
type TemplateMap struct {
	Signup   string // Renders with token when the user signed up.
	Activate string // Renders when the user activated his/her account.
	Auth     string // Renders when the user is logged in.
	Suspend  string // Renders when the user is banned.
}

// PublicServer represents an auth server.
type PublicServer struct {
	PWHashAlgo []crypto.HashFunc
	Templates  *TemplateMap
}

// NewPublicServer creates a new isntance of Server
func NewPublicServer(hashAlgo []crypto.HashFunc) *PublicServer {
	return &PublicServer{PWHashAlgo: hashAlgo}
}
