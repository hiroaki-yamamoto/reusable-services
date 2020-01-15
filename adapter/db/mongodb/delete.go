package mongodb

import "context"

// Delete deletes the matched document by query.
func (me *Mongo) Delete(
	ctx context.Context, query interface{},
) (interface{}, error) {
	return me.col.DeleteOne(ctx, query)
}

// DeleteMany deletes the matched document**s** by query.
func (me *Mongo) DeleteMany(
	ctx context.Context, query interface{},
) (interface{}, error) {
	return me.col.DeleteMany(ctx, query)
}
