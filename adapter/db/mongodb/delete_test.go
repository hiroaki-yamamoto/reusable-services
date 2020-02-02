package mongodb_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
)

var _ = Describe("Delete", func() {
	var chosenDoc *Sample
	BeforeEach(func() {
		chosenDoc = samples[rand.Intn(len(samples))]
	})
	It("Delete should work", func() {
		delCtx, cancelDel := TimeoutContext()
		defer cancelDel()
		deleteCount, err := adapter.Delete(delCtx, chosenDoc)
		Expect(deleteCount).To(BeNumerically("==", 1))
		findCtx, cancelFind := TimeoutContext()
		defer cancelFind()
		num, err := col.CountDocuments(findCtx, chosenDoc)
		Expect(err).To(Succeed())
		Expect(num).To(BeNumerically("==", 0))
	})
})

var _ = Describe("Delete Many", func() {
	var chosenDocs []interface{}
	var chosenID []pr.ObjectID
	BeforeEach(func() {
		chosenDocs = make([]interface{}, 5)
		chosenID = make([]pr.ObjectID, len(chosenDocs))
		for i := range chosenID {
			var id pr.ObjectID
			for ContainObjectID(id, chosenID) {
				doc := samples[rand.Intn(len(samples))]
				chosenDocs[i] = doc
				id = doc.ID
			}
			chosenID[i] = id
		}
	})
	It("DeleteMany should work", func() {
		delCtx, cancelDel := TimeoutContext()
		defer cancelDel()
		deleteCount, err := adapter.DeleteMany(
			delCtx, bson.M{"_id": bson.M{"$in": chosenID}},
		)
		Expect(err).To(Succeed())
		Expect(deleteCount).To(BeNumerically("==", len(chosenDocs)))
		findCtx, cancelFind := TimeoutContext()
		defer cancelFind()
		num, err := col.CountDocuments(
			findCtx, bson.M{"_id": bson.M{"$in": chosenID}},
		)
		Expect(err).To(Succeed())
		Expect(num).To(BeNumerically("==", 0))
	})
})
