package server

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hiroaki-yamamoto/reusable-services/email/rpc"
)

// Send tries to send an email via calling the specified functions.
func (me *Server) Send(
	ctx context.Context,
	req *rpc.SendRequest,
) (*empty.Empty, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var err error
	for _, send := range me.SendFuncs {
		err = send(
			timeoutCtx,
			me.From,
			req.GetEmail(),
			req.GetTitle(),
			req.GetTxtBody(),
			req.GetHtmlBody(),
		)
		if err == nil {
			return &empty.Empty{}, nil
		}
	}
	return nil, err
}
