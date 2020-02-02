package mocks

import "context"
import adp "github.com/hiroaki-yamamoto/reusable-services/adapter"

// MockAdapter Mocks that implements mock.
type MockAdapter struct {
	FindFunc func(
		ctx context.Context,
		query interface{},
		docs interface{},
		opts ...interface{},
	) (err error)
	FindOneFunc func(
		ctx context.Context,
		query interface{},
		doc interface{},
		opts ...interface{},
	) (err error)
	InsertFunc func(
		ctx context.Context, doc interface{},
	) (insertedID interface{}, err error)
	InsertManyFunc func(
		ctx context.Context, docs []interface{},
	) (insertedIDs []interface{}, err error)
	UpdateFunc func(
		ctx context.Context, query interface{}, update interface{},
	) (res *adp.UpdateSummary, err error)
	UpdateManyFunc func(
		ctx context.Context, query interface{}, update interface{},
	) (res *adp.UpdateSummary, err error)
	DeleteFunc func(
		ctx context.Context, filter interface{},
	) (delCount int64, err error)
	DeleteManyFunc func(
		ctx context.Context, filter interface{},
	) (delCount int64, err error)
}

// Find calls FindFunc.
func (me *MockAdapter) Find(
	ctx context.Context,
	query interface{},
	docs interface{},
	opts ...interface{},
) (err error) {
	return me.FindFunc(ctx, query, docs, opts...)
}

// FindOne calls FindOneFunc.
func (me *MockAdapter) FindOne(
	ctx context.Context,
	query interface{},
	doc interface{},
	opts ...interface{},
) (err error) {
	return me.FindOneFunc(ctx, query, doc, opts...)
}

// Insert calls InsertFunc.
func (me *MockAdapter) Insert(
	ctx context.Context, doc interface{},
) (insertedID interface{}, err error) {
	return me.InsertFunc(ctx, doc)
}

// InsertMany calls InsertManyFunc.
func (me *MockAdapter) InsertMany(
	ctx context.Context, docs []interface{},
) (insertedIDs []interface{}, err error) {
	return me.InsertManyFunc(ctx, docs)
}

// Update calls UpdateFunc.
func (me *MockAdapter) Update(
	ctx context.Context, query interface{}, update interface{},
) (res *adp.UpdateSummary, err error) {
	return me.UpdateFunc(ctx, query, update)
}

// UpdateMany calls UpdateManyFunc.
func (me *MockAdapter) UpdateMany(
	ctx context.Context, query interface{}, update interface{},
) (res *adp.UpdateSummary, err error) {
	return me.UpdateManyFunc(ctx, query, update)
}

// Delete calls DeleteFunc.
func (me *MockAdapter) Delete(
	ctx context.Context, filter interface{},
) (delCount int64, err error) {
	return me.DeleteFunc(ctx, filter)
}

// DeleteMany calls DeleteManyFunc.
func (me *MockAdapter) DeleteMany(
	ctx context.Context, filter interface{},
) (delCount int64, err error) {
	return me.DeleteManyFunc(ctx, filter)
}
