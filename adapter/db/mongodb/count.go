package mongodb

import "context"

// Count counts the document.
func (me *Mongo) Count(ctx context.Context, query interface{}) (int64, error) {
	return me.col.CountDocuments(ctx, query)
}
