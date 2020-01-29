package server_test

import (
	"context"
	"testing"
	"time"

	adpMocks "github.com/hiroaki-yamamoto/reusable-services/adapter/mocks"
	"github.com/hiroaki-yamamoto/reusable-services/token/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tokenSize = 32
const tokMaxAge = 48 * time.Hour

var adapter *adpMocks.MockAdapter
var svr *server.Server
var rootCtx context.Context

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = BeforeSuite(func() {
	rootCtx = context.Background()
})

var _ = BeforeEach(func() {
	adapter = &adpMocks.MockAdapter{}
	svr = server.New(adapter, tokenSize, tokMaxAge, 1*time.Second)
})
