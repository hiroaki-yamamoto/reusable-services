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
		curInt, err := adapter.Find(ctx, bson.M{"_id": bson.M{"$in": chosenIDs}})
		Expect(err).To(Succeed())
		cur, ok := curInt.(*mongo.Cursor)
		Expect(ok).To(BeTrue())
		defer cur.Close(ctx)
		ctx, stop := TimeoutContext()
		defer stop()
		var docs []*Sample
		cur.All(ctx, &docs)
		Expect(docs).To(ConsistOf(chosenDocs))
	})
})
