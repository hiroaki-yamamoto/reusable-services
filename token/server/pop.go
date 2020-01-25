package server

import (
	"context"

	"github.com/hiroaki-yamamoto/reusable-services/errors"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
)

// Pop token.
func (me *Server) Pop(
	ctx context.Context, tok *rpc.Token,
) (out *rpc.Token, err error) {
	curCtx, cancel := me.TimeoutContext(ctx)
	defer cancel()
	var tokInt interface{}
	var ok bool
	if tokInt, err = me.adapter.FindOne(curCtx, tok); err == nil {
		if tok, ok = tokInt.(*rpc.Token); !ok {
			err = &errors.InvalidType{}
		}
	}
	return
}
