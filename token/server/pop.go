package server

import (
	"context"

	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
)

// Pop token.
func (me *Server) Pop(
	ctx context.Context, tok *rpc.Token,
) (out *rpc.Token, err error) {
	curCtx, cancel := me.TimeoutContext(ctx)
	defer cancel()
	if err = me.adapter.FindOne(curCtx, map[string]interface{}{
		"purpose": tok.GetPurpose(),
		"token":   tok.GetToken(),
	}, out); err == nil {
		me.adapter.Delete(curCtx, *out)
	}
	return
}
