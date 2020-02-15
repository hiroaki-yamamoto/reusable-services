package server_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	. "github.com/hiroaki-yamamoto/reusable-services/token/server"
)

var _ = Describe("Push", func() {
	var tok *rpc.Token
	BeforeEach(func() {
		tok = &rpc.Token{
			Email:   "test@example.com",
			Purpose: "test",
			Meta:    []byte("Hello world"),
		}
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
		It("Should push the token", func() {
			ctx, stop := context.WithTimeout(rootCtx, 1*time.Second)
			defer stop()
			ret, err := svr.Push(ctx, tok)
			Expect(err).To(Succeed())
			Expect(ret.GetToken()).To(MatchRegexp("^[a-zA-Z0-9]{32}$"))
			tok.Token = ret.GetToken()
			Expect(ret).To(Equal(tok))
		})
	})
	Context("With duplicated token", func() {

	})
})
