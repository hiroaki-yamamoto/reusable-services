package server_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmihailenco/msgpack"

	"github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
)

var _ = Describe("Render", func() {
	Context("With normal", func() {
		It("Should render the template", func() {
			tmpMap := map[string]interface{}{
				"test": "Hello",
			}
			serialized, err := msgpack.Marshal(map[string]interface{}{
				"test": "Hello",
			})
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
})
