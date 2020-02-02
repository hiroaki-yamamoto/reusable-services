package server

import "github.com/hiroaki-yamamoto/reusable-services/token/rpc"

import "time"

// Model represents a model to store a token.
type Model struct {
	*rpc.Token
	Expires time.Time
}
