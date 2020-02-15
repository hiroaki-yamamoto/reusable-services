package server

import (
	"context"
	"errors"

	"github.com/hiroaki-yamamoto/reusable-services/random"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Push issues token id and store & return it.
func (me *Server) Push(ctx context.Context, tok *rpc.Token) (
	ret *rpc.Token, err error,
) {
	if tok == nil {
		err = errors.New("Tok must not be nil")
		return
	}
	processCtx, stop := me.TimeoutContext(ctx)
	defer stop()
	if tok.Token, err = random.GenTxt(32, me.randomTxtSeed); err != nil {
		return
	}
	model := &Model{
		ID:      primitive.NewObjectID(),
		Token:   tok,
		Expires: me.Now().Add(me.maxAge),
	}
	if err = me.Validator.Struct(model); err != nil {
		return
	}
	if len(model.Token.GetToken()) < 1 {
		if err = model.GenerateToken(me.tokSize, me.randomTxtSeed); err != nil {
			return
		}
	}
	if _, err = me.adapter.Insert(processCtx, model); err != nil {
		return
	}
	ret = model.Token
	return
}
