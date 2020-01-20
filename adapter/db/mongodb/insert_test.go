package mongodb_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ = Describe("Insert", func() {
	It("Can insert a single document", func() {
		id := pr.NewObjectID()
		doc := &Sample{ID: id, Number: 0, Meta: "Inserted Doc"}
		ctx, stop := TimeoutContext()
		defer stop()
		resInt, err := adapter.Insert(ctx, doc)
		Expect(err).To(Succeed())
		res, ok := resInt.(*mongo.InsertOneResult)
		Expect(ok).To(BeTrue())
		Expect(res.InsertedID).To(Equal(id))
		findCtx, stopFind := TimeoutContext()
		defer stopFind()
		findRes := col.FindOne(findCtx, bson.M{"_id": id})
		Expect(findRes.Err()).To(Succeed())
		var acDoc *Sample
		Expect(findRes.Decode(&acDoc)).To(Succeed())
		Expect(acDoc).To(Equal(doc))
	})
})

var _ = Describe("InsertMany", func() {
	var docs []*Sample
	var ids []pr.ObjectID
	var toInsert bson.A
	BeforeEach(func() {
		docs = make([]*Sample, 20)
		toInsert = make(bson.A, len(docs))
		ids = make([]pr.ObjectID, len(docs))
		for i := range docs {
			ids[i] = pr.NewObjectID()
			docs[i] = &Sample{
				ID:     ids[i],
				Number: i,
				Meta:   fmt.Sprintf("Inserted %d", i),
			}
			toInsert[i] = docs[i]
		}
	})
	It("Can perform bulk insert", func() {
		ctx, stop := TimeoutContext()
		defer stop()
		resInt, err := adapter.InsertMany(ctx, toInsert)
		Expect(err).To(Succeed())
		res, ok := resInt.(*mongo.InsertManyResult)
		Expect(ok).To(BeTrue())
		Expect(res.InsertedIDs).To(ConsistOf(ids))
		findCtx, stopFind := TimeoutContext()
		defer stopFind()
		cur, err := col.Find(findCtx, bson.M{"_id": bson.M{"$in": ids}})
		Expect(err).Should(Succeed())
		var acdoc []*Sample
		Expect(cur.All(findCtx, &acdoc)).To(Succeed())
		Expect(acdoc).To(ConsistOf(docs))
	})
})
