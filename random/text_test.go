package random_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/hiroaki-yamamoto/reusable-services/random"
)

var _ = Describe("Text", func() {
	It("Shoud generate a random text", func() {
		txts := make([]string, 20)
		const size = 20
		for i := range txts {
			txt := genTxt(size)
			for _, v := range txts[:i] {
				Expect(len(txt)).To(BeEquivalentTo(size))
				Expect(txt).NotTo(Equal(v))
			}
			txts[i] = txt
		}
	})
})
