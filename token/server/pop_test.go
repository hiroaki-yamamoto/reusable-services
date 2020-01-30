package server_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiroaki-yamamoto/reusable-services/random"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
)

var _ = Describe("Pop", func() {
	var deleted bool
	var tokenTxt string
	BeforeEach(func() {
		var err error
		tokenTxt, err = random.GenTxt(tokenSize)
		Expect(err).To(Succeed())
		adapter.FindOneFunc = func(
			ctx context.Context,
			query map[string]interface{},
			doc interface{},
			opts ...interface{},
		) (err error) {
			Expect(query).To(Equal(map[string]interface{}{
				"purpose": "test",
				"token":   tokenTxt,
			}))
			out := doc.(*rpc.Token)
			out.Token = query["token"].(string)
			out.Purpose = query["purpose"].(string)
			return
		}
		adapter.DeleteFunc = func(
			ctx context.Context, doc interface{},
		) (delCount int64, err error) {
			deleted = true
			delCount = 1
			return
		}
	})
	It("Should pop the token", func() {
		res, err := svr.Pop(rootCtx, &rpc.Token{Token: tokenTxt, Purpose: "test"})
		Expect(err).To(Succeed())
		Expect(res).To(Equal(&rpc.Token{Token: tokenTxt, Purpose: "test"}))
		Expect(deleted).To(BeTrue())
	})
})
