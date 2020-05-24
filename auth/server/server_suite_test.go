package server_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	adpMocks "github.com/hiroaki-yamamoto/reusable-services/adapter/mocks"
	"github.com/hiroaki-yamamoto/reusable-services/auth/crypto"
	cryptMocks "github.com/hiroaki-yamamoto/reusable-services/auth/crypto/mocks"
	"github.com/hiroaki-yamamoto/reusable-services/auth/server"
	emailMock "github.com/hiroaki-yamamoto/reusable-services/email/rpc/mocks"
	renderMock "github.com/hiroaki-yamamoto/reusable-services/render/go/rpc/mocks"
	tokMock "github.com/hiroaki-yamamoto/reusable-services/token/rpc/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var rootCtrl *gomock.Controller
var db *adpMocks.MockIAdapter
var emailCli *emailMock.MockEmailClient
var tokCli *tokMock.MockTokenClient
var renderCli *renderMock.MockTemplateServiceClient
var hasher *cryptMocks.MockPasswordHasher
var svr *server.PublicServer

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = BeforeEach(func() {
	rootCtrl = gomock.NewController(GinkgoT())
	db = adpMocks.NewMockIAdapter(rootCtrl)
	hasher = cryptMocks.NewMockPasswordHasher(rootCtrl)
	emailCli = emailMock.NewMockEmailClient(rootCtrl)
	tokCli = tokMock.NewMockTokenClient(rootCtrl)
	renderCli = renderMock.NewMockTemplateServiceClient(rootCtrl)
	svr = server.NewPublicServer(
		db, []crypto.PasswordHasher{hasher},
		zap.NewNop(), &server.TemplateMap{
			Signup: &server.EmailMessage{
				Title:            "Test Signup Title",
				TextTemplateName: "signup.txt",
				HTMLTemplateName: "signup.html",
			},
		},
		emailCli, tokCli, renderCli, "secret",
	)
})

var _ = AfterEach(func() {
	rootCtrl.Finish()
})
