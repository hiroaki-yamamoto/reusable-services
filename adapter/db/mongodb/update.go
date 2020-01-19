package mongodb

import (
	"context"
	"errors"
)

// Update the single matched doc by query, with the diff. "update".
func (me *Mongo) Update(
	ctx context.Context,
	query interface{},
	update interface{},
) (interface{}, error) {
	return nil, errors.New("Not Implemented Yet")
}

// UpdateMany the all matched docs by query, with the diff. "update".
func (me *Mongo) UpdateMany(
	ctx context.Context,
	query interface{},
	update interface{},
) (interface{}, error) {
	return nil, errors.New("Not Implemented Yet")
}
