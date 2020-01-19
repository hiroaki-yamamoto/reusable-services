package mongodb

import (
	"context"
)

// Insert a single document.
func (me *Mongo) Insert(
	ctx context.Context, doc interface{},
) (interface{}, error) {
	return me.col.InsertOne(ctx, doc)
}

// InsertMany store multiple documents into MongoDB.
func (me *Mongo) InsertMany(
	ctx context.Context, docs []interface{},
) (interface{}, error) {
	return me.col.InsertMany(ctx, docs)
}
