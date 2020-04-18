package sendfuncs

// Send is a function type alias to send an email.
type Send func(from string, to string, txtBody, HTMLBody string) error
