package mongodb

import (
	"context"
)

// Find finds documents by the specified query.
func (me *Mongo) Find(
	ctx context.Context,
	query map[string]interface{},
	docs interface{},
	opts ...interface{},
) (err error) {
	if cur, err := me.col.Find(ctx, query); err == nil {
		err = cur.All(ctx, docs)
	}
	return
}

// FindOne finds a document by the specified query
func (me *Mongo) FindOne(
	ctx context.Context,
	query map[string]interface{},
	resdoc interface{},
	opts ...interface{},
) (err error) {
	if sr := me.col.FindOne(ctx, query); sr.Err() == nil {
		err = sr.Decode(resdoc)
	} else {
		err = sr.Err()
	}
	return
}
