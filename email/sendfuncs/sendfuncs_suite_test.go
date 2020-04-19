package sendfuncs_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var rootCtrl *gomock.Controller

func TestSendfuncs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sendfuncs Suite")
}

var _ = BeforeEach(func() {
	rootCtrl = gomock.NewController(GinkgoT())
})

var _ = AfterEach(func() {
	rootCtrl.Finish()
})
