package mongodb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hiroaki-yamamoto/reusable-services/adapter/db/mongodb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client
var db *mongo.Database
var col *mongo.Collection
var adapter *mongodb.Mongo
var rootCtx context.Context

type Sample struct {
	ID     pr.ObjectID `bson:"_id"`
	Number int
	Meta   string
}

var samples []*Sample

// TimeoutContext creates a new timeout context with 3 seconds-timeout from
// rootCtx
func TimeoutContext() (ctx context.Context, cancelFunc context.CancelFunc) {
	ctx, cancelFunc = context.WithTimeout(rootCtx, 3*time.Second)
	return
}

func TestMongodb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mongodb Suite")
}

// ContainObjectID returns true if the specified ObjectID "target" is found
// in a slice of []ObjectID named "arr"
func ContainObjectID(target pr.ObjectID, arr []pr.ObjectID) bool {
	for _, v := range arr {
		if v.Hex() == target.Hex() {
			return true
		}
	}
	return false
}

var _ = BeforeSuite(func() {
	rootCtx = context.Background()
	ctx, cancel := TimeoutContext()
	defer cancel()
	var err error
	cli, err = mongo.Connect(
		ctx, options.Client().ApplyURI("mongodb://adapter:adapter@mongo/"),
	)
	Expect(err).To(Succeed())
	db = cli.Database("adapter-test")
	adapter = mongodb.New(db.Collection("adapters"))
	col = db.Collection("adapters")
})

var _ = BeforeEach(func() {
	samples = make([]*Sample, 20)
	toInsert := make(bson.A, len(samples))
	for i := range samples {
		samples[i] = &Sample{
			ID:     pr.NewObjectID(),
			Number: i,
			Meta:   fmt.Sprintf("Hello %d", i),
		}
		toInsert[i] = samples[i]
	}
	ctx, cancel := TimeoutContext()
	defer cancel()
	col.InsertMany(ctx, toInsert)
})

var _ = AfterEach(func() {
	samples = nil
	ctx, cancel := TimeoutContext()
	defer cancel()
	Expect(col.Drop(ctx)).To(Succeed())
})

var _ = AfterSuite(func() {
	ctx, cancel := TimeoutContext()
	defer cancel()
	Expect(db.Drop(ctx)).To(Succeed())
	Expect(cli.Disconnect(ctx)).To(Succeed())
})
