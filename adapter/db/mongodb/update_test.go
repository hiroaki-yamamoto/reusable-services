package mongodb_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
)

var _ = Describe("Update One", func() {
	var id pr.ObjectID
	var doc Sample
	BeforeEach(func() {
		doc = *samples[rand.Intn(len(samples))]
		doc.Meta = "updated"
		id = doc.ID
	})

	It("Updating the field should be succeeded.", func() {
		ctx, stop := TimeoutContext()
		defer stop()
		q := bson.M{"_id": id}
		res, err := adapter.Update(
			ctx, q, bson.M{"$set": bson.M{"meta": doc.Meta}},
		)
		Expect(err).To(Succeed())
		Expect(res.ModifiedCount).To(BeEquivalentTo(1))
		findCtx, stopFind := TimeoutContext()
		defer stopFind()
		findRes := col.FindOne(findCtx, q)
		Expect(findRes.Err()).To(Succeed())
		var acdocs Sample
		Expect(findRes.Decode(&acdocs)).To(Succeed())
		Expect(acdocs).To(Equal(doc))
	})
})

var _ = Describe("Update Many", func() {
	var chosenID []pr.ObjectID
	var chosenSamples []*Sample
	BeforeEach(func() {
		chosenID = make([]pr.ObjectID, 5)
		chosenSamples = make([]*Sample, 5)
		for i := range chosenID {
			var id pr.ObjectID
			var randomDct Sample
			for ContainObjectID(id, chosenID) {
				randomDct = *samples[rand.Intn(len(samples))]
				id = randomDct.ID
			}
			chosenSamples[i] = &randomDct
			chosenSamples[i].Meta = "updated"
			chosenID[i] = id
		}
	})
	It("Updating the field should be succeeded.", func() {
		ctx, stop := TimeoutContext()
		defer stop()
		q := bson.M{"_id": bson.M{"$in": chosenID}}
		res, err := adapter.UpdateMany(
			ctx, q, bson.M{"$set": bson.M{"meta": "updated"}},
		)
		Expect(res.ModifiedCount).To(BeEquivalentTo(len(chosenID)))
		findCtx, stopFind := TimeoutContext()
		defer stopFind()
		cur, err := col.Find(findCtx, q)
		Expect(err).To(Succeed())
		var acdocs []*Sample
		Expect(cur.All(findCtx, &acdocs)).To(Succeed())
		Expect(acdocs).To(ConsistOf(chosenSamples))
	})
})
