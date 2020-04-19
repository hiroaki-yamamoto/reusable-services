package sendfuncs

import (
	"context"

	"github.com/mailgun/mailgun-go/v4"
	"go.uber.org/zap"
)

// Mailgun indictes a struct to send an email via Mailgun.
type Mailgun struct {
	Logger  *zap.Logger
	Mailgun mailgun.Mailgun
}

// NewMailgun creates a new instance of Mailgun struct.
func NewMailgun(logger *zap.Logger, domain, secret string) *Mailgun {
	return &Mailgun{
		Logger:  logger,
		Mailgun: mailgun.NewMailgun(domain, secret),
	}
}

// Send implements Send Function.
func (me *Mailgun) Send(
	ctx context.Context,
	from, to, subject, txtBody, HTMLBody string,
) error {
	message := me.Mailgun.NewMessage(from, subject, txtBody, to)
	message.SetHtml(HTMLBody)
	status, id, err := me.Mailgun.Send(ctx, message)
	if err != nil {
		return err
	}
	me.Logger.Info(
		"Sending Mail Completed.",
		zap.String("statue", status),
		zap.String("id", id),
	)
	return nil
}
