package server

import (
	"time"

	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"github.com/hiroaki-yamamoto/reusable-services/timeout"
)

// Server is a Token RPC Implementation structure.
type Server struct {
	*timeout.Timeout
	adapter       adapter.IAdapter
	randomTxtSeed []string
	tokSize       int
}

// New creates a new Token RPC Server Instance.
// If randomTxtSeed is empty/nil/zero-value, this service will generate
// random token from ascii letters. i.e. the generated token will be
// matched with [a-zA-Z0-9]{tokSize}
func New(
	adapter adapter.IAdapter,
	tokSize int,
	opTimeout time.Duration,
	randomTxtSeed ...string,
) *Server {
	return &Server{
		Timeout: &timeout.Timeout{
			Timeout: opTimeout,
		},
		adapter:       adapter,
		randomTxtSeed: randomTxtSeed,
		tokSize:       tokSize,
	}
}
