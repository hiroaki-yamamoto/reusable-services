package server

import (
	"context"
	"time"

	"github.com/hiroaki-yamamoto/reusable-services/errors"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	"go.mongodb.org/mongo-driver/bson"
)

// Pop token.
func (me *Server) Pop(
	ctx context.Context, tok *rpc.Token,
) (out *rpc.Token, err error) {
	curCtx, cancel := me.TimeoutContext(ctx)
	defer cancel()
	var res Model
	if err = me.adapter.FindOne(curCtx, bson.M{
		"purpose": tok.GetPurpose(),
		"token":   tok.GetToken(),
	}, &res); err == nil && !res.ID.IsZero() {
		me.adapter.Delete(curCtx, res)
		if res.Expires.After(time.Now().UTC()) {
			out = res.Token
			return
		}
	}
	err = &errors.NotFound{}
	return
}
