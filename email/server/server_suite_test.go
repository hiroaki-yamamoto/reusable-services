package server_test

import (
	"context"
	"testing"

	"github.com/hiroaki-yamamoto/reusable-services/email/sendfuncs"
	"github.com/hiroaki-yamamoto/reusable-services/email/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var svr *server.Server
var callIndexes []int

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = BeforeEach(func() {
	svr = server.New("test@example.com", make([]sendfuncs.Send, 10)...)
	callIndexes = make([]int, 0)
	for i := range svr.SendFuncs {
		svr.SendFuncs[i] = func(index int) sendfuncs.Send {
			return func(
				ctx context.Context, from, to, title, txtBody, HTMLBody string,
			) error {
				callIndexes = append(callIndexes, index)
				return nil
			}
		}(i)
	}
})
