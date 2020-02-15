package server_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hiroaki-yamamoto/reusable-services/random"
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
		var model *Model
		var oldTokTxt string
		BeforeEach(func() {
			var err error
			tok = &rpc.Token{
				Email:   "test@example.com",
				Purpose: "test",
				Meta:    []byte("Hello world"),
			}
			tok.Token, err = random.GenTxt(32)
			Expect(err).To(Succeed())
			oldTokTxt = tok.Token
			model = &Model{
				ID:      primitive.NewObjectID(),
				Token:   tok,
				Expires: now.Add(tokMaxAge),
			}
			adapter.FindOneFunc = func(
				ctx context.Context,
				query interface{},
				doc interface{},
				opts ...interface{},
			) (err error) {
				doc = model
				return
			}
		})
		It("Should update the token", func() {
			ctx, stop := context.WithTimeout(rootCtx, 1*time.Second)
			defer stop()
			ret, err := svr.Push(ctx, tok)
			Expect(err).To(Succeed())
			Expect(ret.GetToken()).To(MatchRegexp("^[a-zA-Z0-9]{32}$"))
			Expect(ret.GetToken()).NotTo(Equal(oldTokTxt))
			tok.Token = ret.GetToken()
			Expect(ret).To(Equal(tok))
		})
	})
})
