package server

import (
	"context"
	"io"
	"strings"

	tmpErrs "github.com/hiroaki-yamamoto/reusable-services/render/go/errors"
	"github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
	"github.com/vmihailenco/msgpack"
)

type itemplate interface {
	Execute(wr io.Writer, data interface{}) error
}

// Render sends the rendered template.
func (me *Server) Render(
	ctx context.Context,
	req *rpc.RenderingRequest,
) (resp *rpc.RenderingResponse, err error) {
	var tmp itemplate
	var ok bool
	tmpName := req.GetTmpName()
	if tmp, ok = me.htmlTemplate[tmpName]; !ok || tmp == nil {
		if tmp, ok = me.textTemplate[tmpName]; !ok || tmp == nil {
			err = &tmpErrs.NoTemplateFound{TmpName: req.GetTmpName()}
			return
		}
	}
	argumentMap := make(map[string]interface{})
	if err = msgpack.Unmarshal(req.GetArgumentMap(), argumentMap); err != nil {
		return
	}
	var buf strings.Builder
	if err = tmp.Execute(&buf, argumentMap); err != nil {
		return
	}
	resp = &rpc.RenderingResponse{Data: buf.String()}
	return
}
