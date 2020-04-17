package server

import (
	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"github.com/hiroaki-yamamoto/reusable-services/auth/crypto"
	"github.com/hiroaki-yamamoto/reusable-services/auth/vldfuncs"
	renderRPC "github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
	tokenRPC "github.com/hiroaki-yamamoto/reusable-services/token/rpc"
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
	TokenCli   tokenRPC.TokenClient
	RenderCli  renderRPC.TemplateServiceClient
	checker    *vld.Validate
}

// NewPublicServer creates a new isntance of Server
func NewPublicServer(
	adapter adapter.IAdapter,
	hashAlgo []crypto.PasswordHasher,
	logger *zap.Logger,
	templates *TemplateMap,
	tokenClient tokenRPC.TokenClient,
	renderClient renderRPC.TemplateServiceClient,
) *PublicServer {
	checker := vld.New()
	checker.RegisterValidationCtx(
		"dbunique", vldfuncs.DBUnique(logger, adapter),
	)
	return &PublicServer{
		PWHashAlgo: hashAlgo,
		Adapter:    adapter,
		Logger:     logger,
		Templates:  templates,
		TokenCli:   tokenClient,
		RenderCli:  renderClient,
		checker:    checker,
	}
}
