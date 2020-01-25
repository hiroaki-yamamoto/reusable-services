package adapter

import "context"

// IAdapter indicates the interface of the adapter.
type IAdapter interface {
	Find(context.Context, interface{}) (interface{}, error)
	FindOne(context.Context, interface{}) (interface{}, error)
	Insert(context.Context, interface{}) (interface{}, error)
	InsertMany(context.Context, []interface{}) (interface{}, error)
	Update(context.Context, interface{}, interface{}) (interface{}, error)
	UpdateMany(context.Context, interface{}, interface{}) (interface{}, error)
	Delete(context.Context, interface{}) (int64, error)
	DeleteMany(context.Context, []interface{}) (int64, error)
}
