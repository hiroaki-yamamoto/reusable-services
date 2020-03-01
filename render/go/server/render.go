package server

import (
	"context"
	"strings"

	errs "github.com/hiroaki-yamamoto/reusable-services/errors"
	"github.com/hiroaki-yamamoto/reusable-services/render/go/interfaces"
	"github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
	"github.com/vmihailenco/msgpack"
)

// Render sends the rendered template.
func (me *Server) Render(
	ctx context.Context,
	req *rpc.RenderingRequest,
) (resp *rpc.RenderingResponse, err error) {
	var tmp interfaces.ITemplate
	var ok bool
	tmpName := req.GetTmpName()
	if tmp, ok = me.templates[tmpName]; !ok || tmp == nil {
		err = &errs.NotFound{
			Metadata: map[string]interface{}{"templateName": tmpName},
		}
		return
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
