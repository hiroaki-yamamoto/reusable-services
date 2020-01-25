package timeout

import "time"

import "context"

// Timeout is a base class to make a context with Timeout.
type Timeout struct {
	Timeout time.Duration
}

// TimeoutContext creates a timeout context.
func (me *Timeout) TimeoutContext(ctx context.Context) (
	context.Context, context.CancelFunc,
) {
	return context.WithTimeout(ctx, me.Timeout)
}
