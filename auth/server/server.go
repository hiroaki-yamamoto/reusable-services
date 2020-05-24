package server

import (
	"sync"

	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"github.com/hiroaki-yamamoto/reusable-services/auth/crypto"
	"github.com/hiroaki-yamamoto/reusable-services/auth/vldfuncs"
	emailRPC "github.com/hiroaki-yamamoto/reusable-services/email/rpc"
	renderRPC "github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
	tokenRPC "github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	"go.uber.org/zap"

	vld "github.com/go-playground/validator/v10"
)

// EmailMessage represents a structure of email message template.
type EmailMessage struct {
	Title            string
	TextTemplateName string
	HTMLTemplateName string
}

// TemplateMap represents a structure that has template names corresponding
// to the templates
type TemplateMap struct {
	Signup   *EmailMessage // Renders with token when the user signed up.
	Activate *EmailMessage // Renders when the user activated his/her account.
	Auth     *EmailMessage // Renders when the user is logged in.
	Suspend  *EmailMessage // Renders when the user is banned.
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
	EmailCli   emailRPC.EmailClient
	checker    *vld.Validate
	WaitGroup  sync.WaitGroup // # of asynchronous jobs.
}

// NewPublicServer creates a new isntance of Server
func NewPublicServer(
	adapter adapter.IAdapter,
	hashAlgo []crypto.PasswordHasher,
	logger *zap.Logger,
	templates *TemplateMap,
	emailClient emailRPC.EmailClient,
	tokenClient tokenRPC.TokenClient,
	renderClient renderRPC.TemplateServiceClient,
	recaptchaSecret string,
) *PublicServer {
	checker := vld.New()
	checker.RegisterValidationCtx(
		"dbunique", vldfuncs.DBUnique(logger, adapter),
	)
	checker.RegisterValidation("base36", vldfuncs.Base36)
	checker.RegisterValidation("recaptcha", vldfuncs.Recaptcha(recaptchaSecret))
	return &PublicServer{
		PWHashAlgo: hashAlgo,
		Adapter:    adapter,
		Logger:     logger,
		Templates:  templates,
		TokenCli:   tokenClient,
		RenderCli:  renderClient,
		EmailCli:   emailClient,
		checker:    checker,
	}
}
