package server

import (
	"strings"
	"time"

	vld "github.com/go-playground/validator/v10"
	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"github.com/hiroaki-yamamoto/reusable-services/timeout"
)

// Server is a Token RPC Implementation structure.
type Server struct {
	*timeout.Timeout
	adapter       adapter.IAdapter
	maxAge        time.Duration
	randomTxtSeed string
	tokSize       int
	Now           func() time.Time
	Validator     *vld.Validate
}

// New creates a new Token RPC Server Instance.
// If randomTxtSeed is empty/nil/zero-value, this service will generate
// random token from ascii letters. i.e. the generated token will be
// matched with [a-zA-Z0-9]{tokSize}
func New(
	adapter adapter.IAdapter,
	tokSize int,
	maxAge time.Duration,
	opTimeout time.Duration,
	randomTxtSeed ...string,
) *Server {
	randTxtSeed :=
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if len(randomTxtSeed) > 0 {
		randTxtSeed = strings.Join(randomTxtSeed, "")
	}
	return &Server{
		Timeout: &timeout.Timeout{
			Timeout: opTimeout,
		},
		Validator:     vld.New(),
		adapter:       adapter,
		maxAge:        maxAge,
		randomTxtSeed: randTxtSeed,
		tokSize:       tokSize,
		Now:           func() time.Time { return time.Now().UTC() },
	}
}
