package server

import (
	"context"

	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
)

// Push issues token id and store & return it.
func (me *Server) Push(ctx context.Context, tok *rpc.Token) (
	ret *rpc.Token, err error,
) {
	return
}
