package mongodb

import (
	"context"
)

// Update the single matched doc by query, with the diff. "update".
func (me *Mongo) Update(
	ctx context.Context,
	query interface{},
	update interface{},
) (interface{}, error) {
	return me.col.UpdateOne(ctx, query, update)
}

// UpdateMany the all matched docs by query, with the diff. "update".
func (me *Mongo) UpdateMany(
	ctx context.Context,
	query interface{},
	update interface{},
) (interface{}, error) {
	return me.col.UpdateMany(ctx, query, update)
}
