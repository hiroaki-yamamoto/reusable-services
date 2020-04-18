package mongodb_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
)

var _ = Describe("Count", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("Should count sample documents.", func() {
		timeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		res, err := adapter.Count(timeout, bson.M{})
		Expect(err).To(Succeed())
		Expect(res).To(BeNumerically("==", len(samples)))
	})
})
