package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hiroaki-yamamoto/reusable-services/email/rpc"
)

// Send tries to send an email via calling the specified functions.
func (me *Server) Send(
	ctx context.Context,
	req *rpc.SendRequest,
) (*empty.Empty, error) {
	var err error
	for _, send := range me.SendFuncs {
		err = send(me.From, req.GetEmail(), req.GetTxtBody(), req.GetHtmlBody())
		if err == nil {
			return &empty.Empty{}, nil
		}
	}
	return nil, err
}
