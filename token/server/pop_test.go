package server_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hiroaki-yamamoto/reusable-services/errors"
	"github.com/hiroaki-yamamoto/reusable-services/random"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	"github.com/hiroaki-yamamoto/reusable-services/token/server"
)

var _ = Describe("Pop", func() {
	var deleted bool
	var tokenTxt string
	BeforeEach(func() {
		adapter.DeleteFunc = func(
			ctx context.Context, query interface{},
		) (delCount int64, err error) {
			deleted = true
			delCount = 1
			return
		}
	})
	Context("For not found token.", func() {
		BeforeEach(func() {
			adapter.FindOneFunc = func(
				ctx context.Context,
				q interface{},
				doc interface{},
				opts ...interface{},
			) (err error) {
				doc = nil
				return
			}
		})
		It("Should Raise NotFound", func() {
			res, err := svr.Pop(rootCtx, &rpc.Token{Token: tokenTxt, Purpose: "test"})
			Expect(res).To(BeNil())
			Expect(err).To(MatchError(&errors.NotFound{}))
			Expect(deleted).To(BeFalse())
		})
	})
	Context("For non-rotted token.", func() {
		BeforeEach(func() {
			var err error
			tokenTxt, err = random.GenTxt(tokenSize)
			Expect(err).To(Succeed())
			adapter.FindOneFunc = func(
				ctx context.Context,
				q interface{},
				doc interface{},
				opts ...interface{},
			) (err error) {
				query, ok := q.(bson.M)
				Expect(ok).To(BeTrue())
				Expect(query).To(Equal(bson.M{
					"purpose": "test",
					"token":   tokenTxt,
				}))
				out := doc.(*server.Model)
				out.ID = pr.NewObjectID()
				out.Token = &rpc.Token{
					Email:   "hello@example.com",
					Token:   query["token"].(string),
					Purpose: query["purpose"].(string),
				}
				out.Expires = time.Now().UTC().Add(2 * time.Hour)
				return
			}
		})
		It("Should pop the token", func() {
			res, err := svr.Pop(rootCtx, &rpc.Token{Token: tokenTxt, Purpose: "test"})
			Expect(err).To(Succeed())
			Expect(res).To(Equal(&rpc.Token{
				Email: "hello@example.com", Token: tokenTxt, Purpose: "test",
			}))
			Expect(deleted).To(BeTrue())
		})
	})
	Context("For rotted token", func() {
		BeforeEach(func() {
			adapter.FindOneFunc = func(
				ctx context.Context,
				q interface{},
				doc interface{},
				opts ...interface{},
			) (err error) {
				query, ok := q.(bson.M)
				Expect(ok).To(BeTrue())
				Expect(query).To(Equal(bson.M{
					"purpose": "test",
					"token":   tokenTxt,
				}))
				out := doc.(*server.Model)
				out.ID = pr.NewObjectID()
				out.Token = &rpc.Token{
					Email:   "hello@example.com",
					Token:   query["token"].(string),
					Purpose: query["purpose"].(string),
				}
				out.Expires = time.Now().UTC().Add(-2 * time.Hour)
				return
			}
		})
		It("Should raise NotFound, and delete the token.", func() {
			res, err := svr.Pop(rootCtx, &rpc.Token{Token: tokenTxt, Purpose: "test"})
			Expect(res).To(BeNil())
			Expect(err).To(MatchError(&errors.NotFound{}))
			Expect(deleted).To(BeTrue())
		})
	})
})
