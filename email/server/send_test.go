package server_test

import (
	"context"
	"fmt"

	"github.com/hiroaki-yamamoto/reusable-services/email/rpc"
	"github.com/hiroaki-yamamoto/reusable-services/email/sendfuncs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Send", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	Context("With Successful Send", func() {
		It("Should call SendFunction once", func() {
			svr.Send(ctx, &rpc.SendRequest{Email: "test2@example.com"})
			Expect(callIndexes).To(Equal([]int{0}))
		})
	})
	Context("With **All** Failure tests", func() {
		BeforeEach(func() {
			for i := range svr.SendFuncs {
				svr.SendFuncs[i] = func(index int) sendfuncs.Send {
					return func(
						from, to, txtBody, HTMLBody string,
					) error {
						callIndexes = append(callIndexes, index)
						return fmt.Errorf("Index: %d", index)
					}
				}(i)
			}
		})
		It("Should raise the corresponding errors.", func() {
			_, err := svr.Send(ctx, &rpc.SendRequest{Email: "test2@example.com"})
			Expect(callIndexes).To(Equal(func() []int {
				ret := make([]int, len(callIndexes))
				for i := range ret {
					ret[i] = i
				}
				return ret
			}()))
			Expect(err).To(MatchError(fmt.Errorf("Index: %d", len(callIndexes)-1)))
		})
	})
	Context("With Partial Failure Tests", func() {
		BeforeEach(func() {
			for i := range svr.SendFuncs {
				if i < 3 {
					svr.SendFuncs[i] = func(index int) sendfuncs.Send {
						return func(
							from, to, txtBody, HTMLBody string,
						) error {
							callIndexes = append(callIndexes, index)
							return fmt.Errorf("Index: %d", index)
						}
					}(i)
				}
			}
		})
		It("Should call the function until succeed.", func() {
			_, err := svr.Send(ctx, &rpc.SendRequest{Email: "test2@example.com"})
			Expect(callIndexes).To(Equal([]int{0, 1, 2, 3}))
			Expect(err).To(Succeed())
		})
	})
})
