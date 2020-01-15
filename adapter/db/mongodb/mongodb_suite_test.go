package mongodb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

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
var rootCtx context.Context

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
	col = db.Collection("adapters")
})

var _ = BeforeEach(func() {
	samples := make(bson.A, 20)
	for i := range samples {
		samples[i] = bson.M{
			"_id":    pr.NewObjectID(),
			"number": i,
			"meta":   fmt.Sprintf("Hello %d", i),
		}
	}
	ctx, cancel := TimeoutContext()
	defer cancel()
	col.InsertMany(ctx, samples)
})

var _ = AfterEach(func() {
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
