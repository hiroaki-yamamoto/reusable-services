package sendfuncs

import "context"

// Send is a function type alias to send an email.
type Send func(
	ctx context.Context,
	from, to, title, txtBody, HTMLBody string,
) error
