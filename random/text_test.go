package random_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/hiroaki-yamamoto/reusable-services/random"
)

var _ = Describe("Text", func() {
	Context("With random seed text", func() {
		const size = 20
		const randomSeedTxt = "abcd"
		var exp string
		BeforeEach(func() {
			exp = fmt.Sprintf("^[a-d]{%d}$", size)
		})
		It("Shoud generate a random text", func() {
			txts := make([]string, size)
			for i := range txts {
				txt, err := GenTxt(size, randomSeedTxt)
				Expect(err).To(Succeed())
				for _, v := range txts[:i] {
					Expect(txt).To(MatchRegexp(exp))
					Expect(txt).NotTo(Equal(v))
				}
				txts[i] = txt
			}
		})
	})
	Context("Without random seed text", func() {
		const size = 20
		It("Shoud generate a random text", func() {
			txts := make([]string, size)
			for i := range txts {
				txt, err := GenTxt(size)
				Expect(err).To(Succeed())
				for _, v := range txts[:i] {
					Expect(txt).To(MatchRegexp(
						fmt.Sprintf("^[a-zA-Z0-9]{%d}$", size)),
					)
					Expect(txt).NotTo(Equal(v))
				}
				txts[i] = txt
			}
		})
	})
})
