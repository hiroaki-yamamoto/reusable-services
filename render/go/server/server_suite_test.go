package server_test

import (
	html "html/template"
	"testing"
	txt "text/template"

	"github.com/hiroaki-yamamoto/reusable-services/render/go/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var svr *server.Server

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = BeforeEach(func() {
	svr = &server.Server{}
	txtTmp := txt.New("txtTmp")
	txtTmp, err := txtTmp.Parse("{{ .test }}")
	Expect(err).To(Succeed())
	htmlTmp := html.New("htmlTmp")
	htmlTmp, err = htmlTmp.Parse("<p>{{ .html }}</p>")
	Expect(err).To(Succeed())

	svr.SetTemplate(txtTmp.Name(), txtTmp)
	svr.SetTemplate(htmlTmp.Name(), htmlTmp)
})
