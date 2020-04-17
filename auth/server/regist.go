package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hiroaki-yamamoto/reusable-services/auth/rpc"
	ms "github.com/mitchellh/mapstructure"
)

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
	// Need to push activation token
}
