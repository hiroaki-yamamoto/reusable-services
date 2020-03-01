package server_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmihailenco/msgpack/v4"

	rsErr "github.com/hiroaki-yamamoto/reusable-services/errors"
	"github.com/hiroaki-yamamoto/reusable-services/render/go/interfaces/mocks"
	"github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
)

type testErr struct{}

func (me *testErr) Error() string {
	return "This is a test error"
}

var _ = Describe("Render", func() {
	Context("With normal", func() {
		It("Should render the template", func() {
			tmpMap := map[string]interface{}{
				"test": "Hello",
			}
			serialized, err := msgpack.Marshal(tmpMap)
			Expect(err).To(Succeed())
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			resp, err := svr.Render(ctx, &rpc.RenderingRequest{
				TmpName: "txtTmp", ArgumentMap: serialized,
			})
			Expect(err).To(Succeed())
			Expect(resp.GetData()).To(Equal(tmpMap["test"]))
		})
	})
	Context("Without template", func() {
		BeforeEach(func() {
			svr.UnsetTemplate("txtTmp")
		})
		It("Should raise NotFound error", func() {
			tmpMap := map[string]interface{}{
				"test": "Hello",
			}
			serialized, err := msgpack.Marshal(tmpMap)
			Expect(err).To(Succeed())
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			resp, err := svr.Render(ctx, &rpc.RenderingRequest{
				TmpName: "txtTmp", ArgumentMap: serialized,
			})
			Expect(err).To(MatchError(&rsErr.NotFound{
				Metadata: map[string]interface{}{"templateName": "txtTmp"},
			}))
			Expect(resp).To(BeNil())
		})
	})
	Context("With invalid rendering arguments", func() {
		It("Should raise an error", func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			resp, err := svr.Render(ctx, &rpc.RenderingRequest{
				TmpName: "txtTmp", ArgumentMap: []byte("はちにんこ"),
			})
			Expect(err).To(MatchError(
				errors.New("msgpack: invalid code=e3 decoding map length"),
			))
			Expect(resp).To(BeNil())
		})
	})
	Context("When the template get an error whle calling Execute", func() {
		var mockCtrl *gomock.Controller
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			iTemplate := mocks.NewMockITemplate(mockCtrl)
			iTemplate.EXPECT().Execute(
				gomock.Any(), gomock.Any(),
			).Return(&testErr{}).Times(1)
			svr.SetTemplate("mock", iTemplate)
		})
		It("Should raise an error", func() {
			tmpMap := map[string]interface{}{
				"test": "Hello",
			}
			serialized, err := msgpack.Marshal(tmpMap)
			Expect(err).To(Succeed())
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			resp, err := svr.Render(ctx, &rpc.RenderingRequest{
				TmpName: "mock", ArgumentMap: serialized,
			})
			Expect(err).To(MatchError(&testErr{}))
			Expect(resp).To(BeNil())
		})
	})
})
