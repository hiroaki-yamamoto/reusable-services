package mongodb_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
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
		var docs []*Sample
		Expect(
			adapter.Find(ctx, bson.M{"_id": bson.M{"$in": chosenIDs}}, &docs),
		).To(Succeed())
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
		var doc *Sample
		Expect(adapter.FindOne(ctx, bson.M{"_id": chosenID}, &doc)).To(Succeed())
		Expect(doc).To(Equal(chosenDoc))
	})
})
