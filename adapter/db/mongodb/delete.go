package mongodb

import (
	"context"
)

// Delete deletes the matched document by query.
func (me *Mongo) Delete(
	ctx context.Context, filter interface{},
) (count int64, err error) {
	if res, err := me.col.DeleteOne(ctx, filter); err == nil {
		count = res.DeletedCount
	}
	return
}

// DeleteMany deletes the matched document**s** by query.
func (me *Mongo) DeleteMany(
	ctx context.Context, filter interface{},
) (count int64, err error) {
	if res, err := me.col.DeleteMany(ctx, filter); err == nil {
		count = res.DeletedCount
	}
	return
}
