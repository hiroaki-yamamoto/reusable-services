package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	adpMocks "github.com/hiroaki-yamamoto/reusable-services/adapter/mocks"
	"github.com/hiroaki-yamamoto/reusable-services/token/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const tokenSize = 32
const tokMaxAge = 48 * time.Hour

var adapter *adpMocks.MockIAdapter
var svr *server.Server
var rootCtx context.Context
var now = time.Now().UTC()
var rootMockCtrl *gomock.Controller

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = BeforeSuite(func() {
	rootMockCtrl = gomock.NewController(GinkgoT())
	rootCtx = context.Background()
})

var _ = BeforeEach(func() {
	adapter = adpMocks.NewMockIAdapter(rootMockCtrl)
	svr = server.New(adapter, tokenSize, tokMaxAge, 1*time.Second)
	svr.Now = func() time.Time { return now }
})
