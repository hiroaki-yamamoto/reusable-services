package mongodb

import (
	"context"
	"errors"
)

// Find finds documents by the specified query.
func (me *Mongo) Find(
	ctx context.Context,
	query interface{},
) (res interface{}, err error) {
	err = errors.New("Not Implemted Yet")
	return
}
