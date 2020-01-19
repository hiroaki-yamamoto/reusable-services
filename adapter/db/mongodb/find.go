package mongodb

import (
	"context"
)

// Find finds documents by the specified query.
func (me *Mongo) Find(
	ctx context.Context,
	query interface{},
) (res interface{}, err error) {
	return me.col.Find(ctx, query)
}

// FindOne finds a document by the specified query
func (me *Mongo) FindOne(
	ctx context.Context,
	query interface{},
) (res interface{}, err error) {
	return me.col.FindOne(ctx, query), nil
}
