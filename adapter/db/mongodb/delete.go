package mongodb

import (
	"context"

	userErr "github.com/hiroaki-yamamoto/reusable-services/errors"
	"go.mongodb.org/mongo-driver/bson"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete deletes the matched document by query.
func (me *Mongo) Delete(
	ctx context.Context, doc interface{},
) (count int64, err error) {
	if res, err := me.col.DeleteOne(ctx, doc); err == nil {
		count = res.DeletedCount
	}
	return
}

// DeleteMany deletes the matched document**s** by query.
func (me *Mongo) DeleteMany(
	ctx context.Context, docs []interface{},
) (count int64, err error) {
	ids := make([]pr.ObjectID, len(docs))
	for i, v := range docs {
		var data []byte
		data, err = bson.Marshal(v)
		if err != nil {
			return
		}
		var doc bson.M
		if err = bson.Unmarshal(data, &doc); err != nil {
			return
		}
		var ok bool
		var id interface{}
		if id, ok = doc["_id"]; !ok {
			err = &userErr.NoIDFound{Value: doc}
			return
		}
		if ids[i], ok = id.(pr.ObjectID); !ok {
			err = &userErr.InvalidType{Value: doc["_id"]}
			return
		}
	}
	if res, err := me.col.DeleteMany(
		ctx, bson.M{"_id": bson.M{"$in": ids}},
	); err == nil {
		count = res.DeletedCount
	}
	return
}
