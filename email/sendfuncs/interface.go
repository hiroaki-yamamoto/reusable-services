package sendfuncs

// Send is a function type alias to send an email.
type Send func(from, to, txtBody, HTMLBody string) error
