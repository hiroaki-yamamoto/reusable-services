package server_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/hiroaki-yamamoto/reusable-services/token/server"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
)

var _ = Describe("Push", func() {
	BeforeEach(func() {
		adapter.FindOneFunc = func(
			ctx context.Context,
			query interface{},
			doc interface{},
			opts ...interface{},
		) (err error) {
			doc = nil
			return
		}
		adapter.InsertFunc = func(
			ctx context.Context, doc interface{},
		) (insertedID interface{}, err error) {
			casted := doc.(*Model)
			return casted.ID, nil
		}
	})
	Context("Without duplicated token", func() {
		ctx, stop := context.WithTimeout(rootCtx, 1 * time.Second)
		defer stop()
		svr.Push(ctx, &rpc.Token{
			Email: "test@example.com",
			Purpose: "test",
			Meta: interface[],
		})
	})
})
