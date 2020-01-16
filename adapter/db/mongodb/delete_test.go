package mongodb_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ = Describe("Delete", func() {
	var chosenID pr.ObjectID
	BeforeEach(func() {
		randomDct, ok := samples[rand.Intn(len(samples))].(bson.M)
		Expect(ok).Should(BeTrue())
		chosenID, ok = randomDct["_id"].(pr.ObjectID)
		Expect(ok).Should(BeTrue())
	})
	It("Delete should work", func() {
		delCtx, cancelDel := TimeoutContext()
		defer cancelDel()
		resInt, err := adapter.Delete(delCtx, bson.M{"_id": chosenID})
		Expect(err).To(Succeed())
		res, ok := resInt.(*mongo.DeleteResult)
		Expect(ok).To(BeTrue())
		Expect(res.DeletedCount).To(BeNumerically("==", 1))
		findCtx, cancelFind := TimeoutContext()
		defer cancelFind()
		num, err := col.CountDocuments(findCtx, bson.M{"_id": chosenID})
		Expect(err).To(Succeed())
		Expect(num).To(BeNumerically("==", 0))
	})
})

var _ = Describe("Delete Many", func() {
	var chosenID []pr.ObjectID
	BeforeEach(func() {
		chosenID = make([]pr.ObjectID, 5)
		for i := range chosenID {
			randomDct, ok := samples[rand.Intn(len(samples))].(bson.M)
			Expect(ok).Should(BeTrue())
			chosenID[i], ok = randomDct["_id"].(pr.ObjectID)
			Expect(ok).Should(BeTrue())
		}
	})
	It("DeleteMany should work", func() {
		delCtx, cancelDel := TimeoutContext()
		defer cancelDel()
		resInt, err := adapter.DeleteMany(
			delCtx, bson.M{"_id": bson.M{"$in": chosenID}},
		)
		Expect(err).To(Succeed())
		res, ok := resInt.(*mongo.DeleteResult)
		Expect(ok).To(BeTrue())
		Expect(res.DeletedCount).To(BeNumerically("==", len(chosenID)))
		findCtx, cancelFind := TimeoutContext()
		defer cancelFind()
		num, err := col.CountDocuments(
			findCtx, bson.M{"_id": bson.M{"$in": chosenID}},
		)
		Expect(err).To(Succeed())
		Expect(num).To(BeNumerically("==", 0))
	})
})
