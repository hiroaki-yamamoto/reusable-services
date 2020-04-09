package server

import "github.com/hiroaki-yamamoto/reusable-services/auth/crypto"

// PublicServer represents an auth server.
type PublicServer struct {
	PWHashAlgo []crypto.HashFunc
}

// NewPublicServer creates a new isntance of Server
func NewPublicServer(hashAlgo []crypto.HashFunc) *PublicServer {
	return &PublicServer{PWHashAlgo: hashAlgo}
}
