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
		ctx context.Context, // The context
		query interface{}, // Query
		docs interface{}, // Pointer to the result document
		opts ...interface{}, // Query options
	) (err error)
	FindOne(
		ctx context.Context, // The context
		query interface{}, // Query
		doc interface{}, // Pointer to the result document
		opts ...interface{}, // Query options
	) (err error)
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
	Delete(ctx context.Context, filter interface{}) (delCount int64, err error)
	DeleteMany(ctx context.Context, filter interface{}) (delCount int64, err error)
}
