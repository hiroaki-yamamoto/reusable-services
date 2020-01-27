package mongodb

import (
	"context"
)

// Find finds documents by the specified query.
func (me *Mongo) Find(
	ctx context.Context,
	query map[string]interface{},
	opts ...interface{},
) (res []interface{}, err error) {
	if cur, err := me.col.Find(ctx, query); err == nil {
		cur.All(ctx, &res)
	}
	return
}

// FindOne finds a document by the specified query
func (me *Mongo) FindOne(
	ctx context.Context,
	query map[string]interface{},
	opts ...interface{},
) (res interface{}, err error) {
	if sr := me.col.FindOne(ctx, query); sr.Err() == nil {
		sr.Decode(&res)
	} else {
		err = sr.Err()
	}
	return
}
