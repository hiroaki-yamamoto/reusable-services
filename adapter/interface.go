package adapter

import "context"

// UpdateSummary represents a structure that tells the result of update/upsert.
type UpdateSummary struct {
	MatchedCount  int64
	ModifiedCount int64
	UpsertedCount int64
	UpsertedIDs   interface{}
}

// IAdapter indicates the interface of the adapter.
type IAdapter interface {
	Find(
		ctx context.Context, query map[string]interface{}, opts ...interface{},
	) (docs []interface{}, err error)
	FindOne(
		ctx context.Context, query map[string]interface{}, opts ...interface{},
	) (doc interface{}, err error)
	Insert(
		ctx context.Context, doc interface{},
	) (insertedID interface{}, err error)
	InsertMany(
		ctx context.Context, docs []interface{},
	) (insertedIDs []interface{}, err error)
	Update(
		ctx context.Context, query interface{}, update interface{},
	) (res *UpdateSummary, err error)
	UpdateMany(
		ctx context.Context, query interface{}, update interface{},
	) (res *UpdateSummary, err error)
	Delete(ctx context.Context, doc interface{}) (delCount int64, err error)
	DeleteMany(ctx context.Context, docs []interface{}) (delCount int64, err error)
}
