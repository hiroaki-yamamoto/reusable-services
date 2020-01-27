package mongodb_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ = Describe("Find", func() {
	var chosenIDs []pr.ObjectID
	var chosenDocs []*Sample
	BeforeEach(func() {
		chosenIDs = make([]pr.ObjectID, 5)
		chosenDocs = make([]*Sample, len(chosenIDs))
		for i := range chosenIDs {
			var chosenDoc *Sample
			var chosenID pr.ObjectID
			for ContainObjectID(chosenID, chosenIDs) {
				chosenDoc = samples[rand.Intn(len(samples))]
				chosenID = chosenDoc.ID
			}
			chosenDocs[i] = chosenDoc
			chosenIDs[i] = chosenID
		}
	})
	It("Should find the correct documents.", func() {
		ctx, cancel := TimeoutContext()
		defer cancel()
		docs, err := adapter.Find(ctx, bson.M{"_id": bson.M{"$in": chosenIDs}})
		Expect(err).To(Succeed())
		Expect(docs).To(ConsistOf(chosenDocs))
	})
})

var _ = Describe("FindOne", func() {
	var chosenID pr.ObjectID
	var chosenDoc *Sample
	BeforeEach(func() {
		chosenDoc = samples[rand.Intn(len(samples))]
		chosenID = chosenDoc.ID
	})
	It("Should find the correct documents.", func() {
		ctx, cancel := TimeoutContext()
		defer cancel()
		curInt, err := adapter.FindOne(ctx, bson.M{"_id": chosenID})
		Expect(err).To(Succeed())
		res, ok := curInt.(*mongo.SingleResult)
		Expect(ok).To(BeTrue())
		Expect(res.Err()).To(Succeed())
		var doc *Sample
		res.Decode(&doc)
		Expect(doc).To(Equal(chosenDoc))
	})
})
