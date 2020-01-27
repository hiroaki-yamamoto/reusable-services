package mongodb

import (
	"context"

	"github.com/hiroaki-yamamoto/reusable-services/adapter"
)

// Update the single matched doc by query, with the diff. "update".
func (me *Mongo) Update(
	ctx context.Context,
	query interface{},
	update interface{},
) (res *adapter.UpdateSummary, err error) {
	if up, err := me.col.UpdateOne(ctx, query, update); err == nil {
		res = &adapter.UpdateSummary{
			MatchedCount:  up.MatchedCount,
			ModifiedCount: up.ModifiedCount,
			UpsertedCount: up.UpsertedCount,
			UpsertedIDs:   up.UpsertedID,
		}
	}
	return
}

// UpdateMany the all matched docs by query, with the diff. "update".
func (me *Mongo) UpdateMany(
	ctx context.Context,
	query interface{},
	update interface{},
) (res *adapter.UpdateSummary, err error) {
	if up, err := me.col.UpdateMany(ctx, query, update); err == nil {
		res = &adapter.UpdateSummary{
			MatchedCount:  up.MatchedCount,
			ModifiedCount: up.ModifiedCount,
			UpsertedCount: up.UpsertedCount,
			UpsertedIDs:   up.UpsertedID,
		}
	}
	return
}
