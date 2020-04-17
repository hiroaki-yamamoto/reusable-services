package server

import (
	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"github.com/hiroaki-yamamoto/reusable-services/auth/crypto"
	"github.com/hiroaki-yamamoto/reusable-services/auth/vldfuncs"
	"go.uber.org/zap"

	vld "github.com/go-playground/validator/v10"
)

// TemplateMap represents a structure that has template names corresponding
// to the templates
type TemplateMap struct {
	Signup   string // Renders with token when the user signed up.
	Activate string // Renders when the user activated his/her account.
	Auth     string // Renders when the user is logged in.
	Suspend  string // Renders when the user is banned.
}

// PublicServer represents an auth server.
// Note that this server depedns on token and render services.
type PublicServer struct {
	Adapter    adapter.IAdapter
	PWHashAlgo []crypto.PasswordHasher
	Logger     *zap.Logger
	Templates  *TemplateMap
	checker    *vld.Validate
}

// NewPublicServer creates a new isntance of Server
func NewPublicServer(
	hashAlgo []crypto.PasswordHasher,
	adapter adapter.IAdapter,
	logger *zap.Logger,
) *PublicServer {
	checker := vld.New()
	checker.RegisterValidationCtx(
		"dbunique", vldfuncs.DBUnique(logger, adapter),
	)
	return &PublicServer{
		PWHashAlgo: hashAlgo,
		Adapter:    adapter,
		Logger:     logger,
		checker:    checker,
	}
}
