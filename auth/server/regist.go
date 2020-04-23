package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hiroaki-yamamoto/reusable-services/auth/rpc"
	emailRPC "github.com/hiroaki-yamamoto/reusable-services/email/rpc"
	renderRPC "github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
	tokenRPC "github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	ms "github.com/mitchellh/mapstructure"
	"github.com/vmihailenco/msgpack/v4"
	"go.uber.org/zap"
)

type renderRespError struct {
	Resp *renderRPC.RenderingResponse
	Err  error
}

// SignUp implements Signup RPC to register the user.
// Note that registering the user doesn't mean activate the user. Activation
// process should be required by end-user.
func (me *PublicServer) SignUp(
	ctx context.Context,
	req *rpc.RegistRequest,
) (*empty.Empty, error) {
	model := &Auth{}
	if err := ms.Decode(req, model); err != nil {
		return nil, err
	}
	model.State = Inactive
	if err := me.checker.StructPartialCtx(
		ctx, model, "Email", "State",
	); err != nil {
		return nil, err
	}
	if _, err := me.Adapter.Insert(ctx, model); err != nil {
		return nil, err
	}
	tok, err := me.TokenCli.Push(ctx, &tokenRPC.Token{
		Email:   req.GetEmail(),
		Purpose: "activation",
	})
	if err != nil {
		return nil, err
	}
	kwarg, err := msgpack.Marshal(
		map[string]interface{}{"token": tok.GetToken()},
	)
	if err != nil {
		return nil, err
	}
	txtResChan := make(chan *renderRespError)
	htmlResChan := make(chan *renderRespError)
	me.WaitGroup.Add(2)
	go func() {
		defer me.WaitGroup.Done()
		res, err := me.RenderCli.Render(ctx, &renderRPC.RenderingRequest{
			TmpName:     me.Templates.Text.Signup,
			ArgumentMap: kwarg,
		})
		txtResChan <- &renderRespError{Resp: res, Err: err}
	}()
	go func() {
		defer me.WaitGroup.Done()
		res, err := me.RenderCli.Render(ctx, &renderRPC.RenderingRequest{
			TmpName:     me.Templates.HTML.Signup,
			ArgumentMap: kwarg,
		})
		htmlResChan <- &renderRespError{Resp: res, Err: err}
	}()
	txtRes := <-txtResChan
	htmlRes := <-htmlResChan
	if txtRes.Err != nil {
		return nil, txtRes.Err
	}
	if htmlRes.Err != nil {
		return nil, htmlRes.Err
	}
	me.WaitGroup.Add(1)
	go func() {
		defer me.WaitGroup.Done()
		emailReq := &emailRPC.SendRequest{
			Email:    model.Email,
			Title:    me.Templates.Title,
			TxtBody:  txtRes.Resp.GetData(),
			HtmlBody: htmlRes.Resp.GetData(),
		}
		_, err = me.EmailCli.Send(ctx, emailReq)
		if err != nil {
			me.Logger.Error(
				"An error has been occured on sending an email:",
				zap.Any("req", emailReq),
				zap.Error(err),
			)
		}
	}()
	return &empty.Empty{}, nil
}
