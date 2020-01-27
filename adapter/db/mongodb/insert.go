package mongodb

import (
	"context"
)

// Insert a single document.
func (me *Mongo) Insert(
	ctx context.Context, doc interface{},
) (insID interface{}, err error) {
	if insRes, err := me.col.InsertOne(ctx, doc); err == nil {
		insID = insRes.InsertedID
	}
	return
}

// InsertMany store multiple documents into MongoDB.
func (me *Mongo) InsertMany(
	ctx context.Context, docs []interface{},
) (insertedIDs []interface{}, err error) {
	if res, err := me.col.InsertMany(ctx, docs); err == nil {
		insertedIDs = res.InsertedIDs
	}
	return
}
